// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package dataplex_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccDataplexEntryGroupIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexEntryGroupIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccDataplexEntryGroupIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccDataplexEntryGroupIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccDataplexEntryGroupIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccDataplexEntryGroupIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexEntryGroupIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_dataplex_entry_group_iam_policy.foo", "policy_data"),
			},
			{
				Config: testAccDataplexEntryGroupIamPolicy_emptyBinding(context),
			},
		},
	})
}

func testAccDataplexEntryGroupIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_entry_group" "test_entry_group_basic" {
  entry_group_id = "tf-test-entry-group-basic%{random_suffix}"
  project = "%{project_name}"
  location = "us-central1"
}

resource "google_dataplex_entry_group_iam_member" "foo" {
  project = google_dataplex_entry_group.test_entry_group_basic.project
  location = google_dataplex_entry_group.test_entry_group_basic.location
  entry_group_id = google_dataplex_entry_group.test_entry_group_basic.entry_group_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccDataplexEntryGroupIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_entry_group" "test_entry_group_basic" {
  entry_group_id = "tf-test-entry-group-basic%{random_suffix}"
  project = "%{project_name}"
  location = "us-central1"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_dataplex_entry_group_iam_policy" "foo" {
  project = google_dataplex_entry_group.test_entry_group_basic.project
  location = google_dataplex_entry_group.test_entry_group_basic.location
  entry_group_id = google_dataplex_entry_group.test_entry_group_basic.entry_group_id
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_dataplex_entry_group_iam_policy" "foo" {
  project = google_dataplex_entry_group.test_entry_group_basic.project
  location = google_dataplex_entry_group.test_entry_group_basic.location
  entry_group_id = google_dataplex_entry_group.test_entry_group_basic.entry_group_id
  depends_on = [
    google_dataplex_entry_group_iam_policy.foo
  ]
}
`, context)
}

func testAccDataplexEntryGroupIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_entry_group" "test_entry_group_basic" {
  entry_group_id = "tf-test-entry-group-basic%{random_suffix}"
  project = "%{project_name}"
  location = "us-central1"
}

data "google_iam_policy" "foo" {
}

resource "google_dataplex_entry_group_iam_policy" "foo" {
  project = google_dataplex_entry_group.test_entry_group_basic.project
  location = google_dataplex_entry_group.test_entry_group_basic.location
  entry_group_id = google_dataplex_entry_group.test_entry_group_basic.entry_group_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataplexEntryGroupIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_entry_group" "test_entry_group_basic" {
  entry_group_id = "tf-test-entry-group-basic%{random_suffix}"
  project = "%{project_name}"
  location = "us-central1"
}

resource "google_dataplex_entry_group_iam_binding" "foo" {
  project = google_dataplex_entry_group.test_entry_group_basic.project
  location = google_dataplex_entry_group.test_entry_group_basic.location
  entry_group_id = google_dataplex_entry_group.test_entry_group_basic.entry_group_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccDataplexEntryGroupIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_entry_group" "test_entry_group_basic" {
  entry_group_id = "tf-test-entry-group-basic%{random_suffix}"
  project = "%{project_name}"
  location = "us-central1"
}

resource "google_dataplex_entry_group_iam_binding" "foo" {
  project = google_dataplex_entry_group.test_entry_group_basic.project
  location = google_dataplex_entry_group.test_entry_group_basic.location
  entry_group_id = google_dataplex_entry_group.test_entry_group_basic.entry_group_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}