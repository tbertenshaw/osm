package catalog

import (
	"fmt"
	"time"

	testclient "k8s.io/client-go/kubernetes/fake"

	set "github.com/deckarep/golang-set"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/open-service-mesh/osm/pkg/certificate"
	"github.com/open-service-mesh/osm/pkg/certificate/providers/tresor"
	"github.com/open-service-mesh/osm/pkg/endpoint"
	"github.com/open-service-mesh/osm/pkg/ingress"
	"github.com/open-service-mesh/osm/pkg/providers/kube"
	"github.com/open-service-mesh/osm/pkg/service"
	"github.com/open-service-mesh/osm/pkg/smi"
	"github.com/open-service-mesh/osm/pkg/tests"
	"github.com/open-service-mesh/osm/pkg/trafficpolicy"
)

var _ = Describe("Catalog tests", func() {
	endpointProviders := []endpoint.Provider{kube.NewFakeProvider()}
	kubeClient := testclient.NewSimpleClientset()
	cache := make(map[certificate.CommonName]certificate.Certificater)
	certManager := tresor.NewFakeCertManager(&cache, 1*time.Hour)
	meshCatalog := NewMeshCatalog(kubeClient, smi.NewFakeMeshSpecClient(), certManager, ingress.NewFakeIngressMonitor(), make(<-chan struct{}), endpointProviders...)

	Context("Test ListTrafficPolicies", func() {
		It("lists traffic policies", func() {
			actual, err := meshCatalog.ListTrafficPolicies(tests.BookstoreService)
			Expect(err).ToNot(HaveOccurred())

			expected := []trafficpolicy.TrafficTarget{tests.TrafficPolicy}
			Expect(actual).To(Equal(expected))
		})
	})

	Context("Test getActiveServices", func() {
		It("lists active services", func() {
			actual := meshCatalog.getActiveServices(set.NewSet(tests.BookstoreService))
			expected := set.NewSet(service.NamespacedService{
				Namespace: "default",
				Service:   "bookstore",
			})
			Expect(actual.Equal(expected)).To(Equal(true))
		})
	})

	Context("Test getTrafficPolicyPerRoute", func() {
		It("lists traffic policies", func() {
			allTrafficPolicies, err := getTrafficPolicyPerRoute(meshCatalog, tests.RoutePolicyMap, tests.BookstoreService)
			Expect(err).ToNot(HaveOccurred())

			expected := []trafficpolicy.TrafficTarget{{
				Name: tests.TrafficTargetName,
				Destination: trafficpolicy.TrafficResource{
					ServiceAccount: tests.BookstoreServiceAccountName,
					Namespace:      tests.Namespace,
					Services:       set.NewSet(tests.BookstoreService),
				},
				Source: trafficpolicy.TrafficResource{
					ServiceAccount: tests.BookbuyerServiceAccountName,
					Namespace:      tests.Namespace,
					Services:       set.NewSet(tests.BookbuyerService),
				},
				Routes: []trafficpolicy.Route{{PathRegex: "", Methods: nil}},
			}}

			Expect(allTrafficPolicies).To(Equal(expected))
		})
	})

	Context("Test listServicesForServiceAccount", func() {
		mc := MeshCatalog{
			serviceAccountToServicesCache: map[service.NamespacedServiceAccount][]service.NamespacedService{
				tests.BookstoreServiceAccount: {tests.BookstoreService},
			},
		}
		It("lists services for service account", func() {
			actual, err := mc.listServicesForServiceAccount(tests.BookstoreServiceAccount)
			Expect(err).ToNot(HaveOccurred())
			expected := set.NewSet(service.NamespacedService{
				Namespace: tests.Namespace,
				Service:   tests.BookstoreServiceName,
			})
			Expect(actual.Equal(expected)).To(Equal(true))
		})
	})

	Context("Test getHTTPPathsPerRoute", func() {
		mc := MeshCatalog{meshSpec: smi.NewFakeMeshSpecClient()}
		It("constructs HTTP paths per route", func() {
			actual, err := mc.getHTTPPathsPerRoute()
			Expect(err).ToNot(HaveOccurred())

			key := fmt.Sprintf("HTTPRouteGroup/%s/%s/%s", tests.Namespace, tests.RouteGroupName, tests.MatchName)
			expected := map[string]trafficpolicy.Route{
				key: {
					PathRegex: tests.BookstoreBuyPath,
					Methods:   []string{"GET"},
				},
			}
			Expect(actual).To(Equal(expected))
		})
	})
})
