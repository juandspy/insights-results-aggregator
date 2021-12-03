/*
Copyright © 2021 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tests

import (
	httputils "github.com/RedHatInsights/insights-operator-utils/http"
	"github.com/RedHatInsights/insights-results-aggregator-data/testdata"
	"github.com/verdverm/frisby"

	"github.com/RedHatInsights/insights-results-aggregator/server"
	"github.com/RedHatInsights/insights-results-aggregator/types"
)

// constructURLForReportInfoForOrgCluster function constructs an URL to access
// the endpoint to retrieve results metadata for given cluster from selected
// organization
func constructURLForReportInfoForOrgCluster(organizationID string,
	clusterID string, userID types.UserID) string {
	return httputils.MakeURLToEndpoint(apiURL,
		server.ReportMetainfoEndpoint, organizationID, clusterID, userID)
}

func constructFrisby(msg, url string, statusCode int, setAuth bool) *frisby.Frisby {
	f := frisby.Create(msg).Get(url)
	if setAuth {
		setAuthHeader(f)
	}
	f.Send()
	f.ExpectStatus(statusCode)
	f.ExpectHeader(contentTypeHeader, ContentTypeJSON)
	return f
}

// checkReportInfoEndpointForKnownOrganizationAndKnownCluster check if the
// endpoint to return report metadata works as expected
func checkReportInfoEndpointForKnownOrganizationAndKnownCluster() {
	url := constructURLForReportInfoForOrgCluster("1", knownClusterForOrganization1, testdata.UserID)
	f := constructFrisby("Check the endpoint to return report metadata for existing organization and cluster ID", url, 200, true)
	f.PrintReport()
}

// checkReportInfoEndpointForKnownOrganizationAndUnknownCluster check if the
// endpoint to return report metadata works as expected
func checkReportInfoEndpointForKnownOrganizationAndUnknownCluster() {
	url := constructURLForReportInfoForOrgCluster("1", unknownClusterForOrganization1, testdata.UserID)
	f := constructFrisby("Check the endpoint to return report metadata for existing organization and non-existing cluster ID", url, 404, true)
	f.PrintReport()
}

// checkReportInfoEndpointForUnknownOrganizationAndKnownCluster check if the
// endpoint to return report metadata works as expected
func checkReportInfoEndpointForUnknownOrganizationAndKnownCluster() {
	url := constructURLForReportInfoForOrgCluster(unknownOrganizationID, knownClusterForOrganization1, testdata.UserID)
	f := constructFrisby("Check the endpoint to return report metadata for unknown organization and cluster ID", url, 404, true)
	f.PrintReport()
}

// checkReportInfoEndpointForUnknownOrganizationAndUnknownCluster check if the
// endpoint to return report metadata works as expected
func checkReportInfoEndpointForUnknownOrganizationAndUnknownCluster() {
	url := constructURLForReportInfoForOrgCluster(unknownOrganizationID, unknownClusterForOrganization1, testdata.UserID)
	f := constructFrisby("Check the endpoint to return report metadata for unknown organization and non-existing cluster ID", url, 404, true)
	f.PrintReport()
}

// checkReportInfoEndpointForImproperOrganization check if the endpoint to
// return report metadata works as expected
func checkReportInfoEndpointForImproperOrganization() {
	url := constructURLForReportInfoForOrgCluster(wrongOrganizationID, knownClusterForOrganization1, testdata.UserID)
	f := constructFrisby("Check the endpoint to return report metadata for improper organization", url, 400, true)

	statusResponse := readStatusFromResponse(f)
	checkErrorStatusResponse(f, statusResponse)

	f.PrintReport()
}

// checkReportInfoEndpointWrongMethods check if the endpoint to return results
// responds correctly to other methods than HTTP GET
func checkReportInfoEndpointWrongMethods() {
	url := constructURLForReportInfoForOrgCluster("1", knownClusterForOrganization1, testdata.UserID)
	checkGetEndpointByOtherMethods(url, false)
}

// checkReportInfoEndpointForKnownOrganizationAndKnownClusterUnauthorizedCase
// check if the endpoint to return report metadata works as expected.
// This test variant does not sent authorization header
func checkReportInfoEndpointForKnownOrganizationAndKnownClusterUnauthorizedCase() {
	url := constructURLForReportInfoForOrgCluster("1", knownClusterForOrganization1, testdata.UserID)
	f := constructFrisby("Check the endpoint to return report metadata for existing organization and cluster ID", url, 401, false)
	f.PrintReport()
}

// checkReportInfoEndpointForKnownOrganizationAndUnknownClusterUnauthorizedCase
// check if the endpoint to return report metadata works as expected.
// This test variant does not sent authorization header
func checkReportInfoEndpointForKnownOrganizationAndUnknownClusterUnauthorizedCase() {
	url := constructURLForReportInfoForOrgCluster("1", unknownClusterForOrganization1, testdata.UserID)
	f := constructFrisby("Check the endpoint to return report metadata for existing organization and non-existing cluster ID", url, 401, false)
	f.PrintReport()
}

// checkReportInfoEndpointForUnknownOrganizationAndKnownClusterUnauthorizedCase
// check if the endpoint to return report metadata works as expected.
// This test variant does not sent authorization header
func checkReportInfoEndpointForUnknownOrganizationAndKnownClusterUnauthorizedCase() {
	url := constructURLForReportInfoForOrgCluster(unknownOrganizationID, knownClusterForOrganization1, testdata.UserID)
	f := constructFrisby("Check the endpoint to return report metadata for unknown organization and cluster ID", url, 401, false)
	f.PrintReport()
}

// checkReportInfoEndpointForUnknownOrganizationAndUnknownClusterUnauthorizedCase
// check if the endpoint to return report metadata works as expected.
// This test variant does not sent authorization header
func checkReportInfoEndpointForUnknownOrganizationAndUnknownClusterUnauthorizedCase() {
	url := constructURLForReportInfoForOrgCluster(unknownOrganizationID, unknownClusterForOrganization1, testdata.UserID)
	f := constructFrisby("Check the endpoint to return report metadata for unknown organization and non-existing cluster ID", url, 401, false)
	f.PrintReport()
}

// checkReportInfoEndpointForImproperOrganizationUnauthorizedCase check if the
// endpoint to return report metadata works as expected.
// This test variant does not sent authorization header
func checkReportInfoEndpointForImproperOrganizationUnauthorizedCase() {
	url := constructURLForReportInfoForOrgCluster(wrongOrganizationID, knownClusterForOrganization1, testdata.UserID)
	f := constructFrisby("Check the endpoint to return report metadata for improper organization", url, 401, false)

	statusResponse := readStatusFromResponse(f)
	checkErrorStatusResponse(f, statusResponse)

	f.PrintReport()
}
