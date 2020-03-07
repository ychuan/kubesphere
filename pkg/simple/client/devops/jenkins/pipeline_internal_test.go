package jenkins

import (
	"kubesphere.io/kubesphere/pkg/simple/client/devops"
	"reflect"
	"testing"
)

func Test_NoScmPipelineConfig(t *testing.T) {
	inputs := []*devops.NoScmPipeline{
		{
			Name:        "",
			Description: "for test",
			Jenkinsfile: "node{echo 'hello'}",
		},
		{
			Name:        "",
			Description: "",
			Jenkinsfile: "node{echo 'hello'}",
		},
		{
			Name:              "",
			Description:       "",
			Jenkinsfile:       "node{echo 'hello'}",
			DisableConcurrent: true,
		},
	}
	for _, input := range inputs {
		outputString, err := createPipelineConfigXml(input)
		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		output, err := parsePipelineConfigXml(outputString)

		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		if !reflect.DeepEqual(input, output) {
			t.Fatalf("input [%+v] output [%+v] should equal ", input, output)
		}
	}
}

func Test_NoScmPipelineConfig_Discarder(t *testing.T) {
	inputs := []*devops.NoScmPipeline{
		{
			Name:        "",
			Description: "for test",
			Jenkinsfile: "node{echo 'hello'}",
			Discarder: &devops.DiscarderProperty{
				DaysToKeep: "3", NumToKeep: "5",
			},
		},
		{
			Name:        "",
			Description: "for test",
			Jenkinsfile: "node{echo 'hello'}",
			Discarder: &devops.DiscarderProperty{
				DaysToKeep: "3", NumToKeep: "",
			},
		},
		{
			Name:        "",
			Description: "for test",
			Jenkinsfile: "node{echo 'hello'}",
			Discarder: &devops.DiscarderProperty{
				DaysToKeep: "", NumToKeep: "21321",
			},
		},
		{
			Name:        "",
			Description: "for test",
			Jenkinsfile: "node{echo 'hello'}",
			Discarder: &devops.DiscarderProperty{
				DaysToKeep: "", NumToKeep: "",
			},
		},
	}
	for _, input := range inputs {
		outputString, err := createPipelineConfigXml(input)
		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		output, err := parsePipelineConfigXml(outputString)

		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		if !reflect.DeepEqual(input, output) {
			t.Fatalf("input [%+v] output [%+v] should equal ", input, output)
		}
	}
}

func Test_NoScmPipelineConfig_Param(t *testing.T) {
	inputs := []*devops.NoScmPipeline{
		{
			Name:        "",
			Description: "for test",
			Jenkinsfile: "node{echo 'hello'}",
			Parameters: &devops.Parameters{
				&devops.Parameter{
					Name:         "d",
					DefaultValue: "a\nb",
					Type:         "choice",
					Description:  "fortest",
				},
			},
		},
		{
			Name:        "",
			Description: "for test",
			Jenkinsfile: "node{echo 'hello'}",
			Parameters: &devops.Parameters{
				&devops.Parameter{
					Name:         "a",
					DefaultValue: "abc",
					Type:         "string",
					Description:  "fortest",
				},
				&devops.Parameter{
					Name:         "b",
					DefaultValue: "false",
					Type:         "boolean",
					Description:  "fortest",
				},
				&devops.Parameter{
					Name:         "c",
					DefaultValue: "password \n aaa",
					Type:         "text",
					Description:  "fortest",
				},
				&devops.Parameter{
					Name:         "d",
					DefaultValue: "a\nb",
					Type:         "choice",
					Description:  "fortest",
				},
			},
		},
	}
	for _, input := range inputs {
		outputString, err := createPipelineConfigXml(input)
		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		output, err := parsePipelineConfigXml(outputString)

		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		if !reflect.DeepEqual(input, output) {
			t.Fatalf("input [%+v] output [%+v] should equal ", input, output)
		}
	}
}

func Test_NoScmPipelineConfig_Trigger(t *testing.T) {
	inputs := []*devops.NoScmPipeline{
		{
			Name:        "",
			Description: "for test",
			Jenkinsfile: "node{echo 'hello'}",
			TimerTrigger: &devops.TimerTrigger{
				Cron: "1 1 1 * * *",
			},
		},

		{
			Name:        "",
			Description: "for test",
			Jenkinsfile: "node{echo 'hello'}",
			RemoteTrigger: &devops.RemoteTrigger{
				Token: "abc",
			},
		},
		{
			Name:        "",
			Description: "for test",
			Jenkinsfile: "node{echo 'hello'}",
			TimerTrigger: &devops.TimerTrigger{
				Cron: "1 1 1 * * *",
			},
			RemoteTrigger: &devops.RemoteTrigger{
				Token: "abc",
			},
		},
	}

	for _, input := range inputs {
		outputString, err := createPipelineConfigXml(input)
		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		output, err := parsePipelineConfigXml(outputString)

		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		if !reflect.DeepEqual(input, output) {
			t.Fatalf("input [%+v] output [%+v] should equal ", input, output)
		}
	}
}

func Test_MultiBranchPipelineConfig(t *testing.T) {

	inputs := []*devops.MultiBranchPipeline{
		{
			Name:        "",
			Description: "for test",
			ScriptPath:  "Jenkinsfile",
			SourceType:  "git",
			GitSource:   &devops.GitSource{},
		},
		{
			Name:         "",
			Description:  "for test",
			ScriptPath:   "Jenkinsfile",
			SourceType:   "github",
			GitHubSource: &devops.GithubSource{},
		},
		{
			Name:            "",
			Description:     "for test",
			ScriptPath:      "Jenkinsfile",
			SourceType:      "single_svn",
			SingleSvnSource: &devops.SingleSvnSource{},
		},
		{
			Name:        "",
			Description: "for test",
			ScriptPath:  "Jenkinsfile",
			SourceType:  "svn",
			SvnSource:   &devops.SvnSource{},
		},
	}
	for _, input := range inputs {
		outputString, err := createMultiBranchPipelineConfigXml("", input)
		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		output, err := parseMultiBranchPipelineConfigXml(outputString)

		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		if !reflect.DeepEqual(input, output) {
			t.Fatalf("input [%+v] output [%+v] should equal ", input, output)
		}
	}
}

func Test_MultiBranchPipelineConfig_Discarder(t *testing.T) {

	inputs := []*devops.MultiBranchPipeline{
		{
			Name:        "",
			Description: "for test",
			ScriptPath:  "Jenkinsfile",
			SourceType:  "git",
			Discarder: &devops.DiscarderProperty{
				DaysToKeep: "1",
				NumToKeep:  "2",
			},
			GitSource: &devops.GitSource{},
		},
	}
	for _, input := range inputs {
		outputString, err := createMultiBranchPipelineConfigXml("", input)
		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		output, err := parseMultiBranchPipelineConfigXml(outputString)

		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		if !reflect.DeepEqual(input, output) {
			t.Fatalf("input [%+v] output [%+v] should equal ", input, output)
		}
	}
}

func Test_MultiBranchPipelineConfig_TimerTrigger(t *testing.T) {
	inputs := []*devops.MultiBranchPipeline{
		{
			Name:        "",
			Description: "for test",
			ScriptPath:  "Jenkinsfile",
			SourceType:  "git",
			TimerTrigger: &devops.TimerTrigger{
				Interval: "12345566",
			},
			GitSource: &devops.GitSource{},
		},
	}
	for _, input := range inputs {
		outputString, err := createMultiBranchPipelineConfigXml("", input)
		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		output, err := parseMultiBranchPipelineConfigXml(outputString)

		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		if !reflect.DeepEqual(input, output) {
			t.Fatalf("input [%+v] output [%+v] should equal ", input, output)
		}
	}
}

func Test_MultiBranchPipelineConfig_Source(t *testing.T) {

	inputs := []*devops.MultiBranchPipeline{
		{
			Name:        "",
			Description: "for test",
			ScriptPath:  "Jenkinsfile",
			SourceType:  "git",
			TimerTrigger: &devops.TimerTrigger{
				Interval: "12345566",
			},
			GitSource: &devops.GitSource{
				Url:              "https://github.com/kubesphere/devops",
				CredentialId:     "git",
				DiscoverBranches: true,
			},
		},
		{
			Name:        "",
			Description: "for test",
			ScriptPath:  "Jenkinsfile",
			SourceType:  "github",
			TimerTrigger: &devops.TimerTrigger{
				Interval: "12345566",
			},
			GitHubSource: &devops.GithubSource{
				Owner:                "kubesphere",
				Repo:                 "devops",
				CredentialId:         "github",
				ApiUri:               "https://api.github.com",
				DiscoverBranches:     1,
				DiscoverPRFromOrigin: 2,
				DiscoverPRFromForks: &devops.DiscoverPRFromForks{
					Strategy: 1,
					Trust:    1,
				},
			},
		},
		{
			Name:        "",
			Description: "for test",
			ScriptPath:  "Jenkinsfile",
			SourceType:  "bitbucket_server",
			TimerTrigger: &devops.TimerTrigger{
				Interval: "12345566",
			},
			BitbucketServerSource: &devops.BitbucketServerSource{
				Owner:                "kubesphere",
				Repo:                 "devops",
				CredentialId:         "github",
				ApiUri:               "https://api.github.com",
				DiscoverBranches:     1,
				DiscoverPRFromOrigin: 2,
				DiscoverPRFromForks: &devops.DiscoverPRFromForks{
					Strategy: 1,
					Trust:    1,
				},
			},
		},

		{
			Name:        "",
			Description: "for test",
			ScriptPath:  "Jenkinsfile",
			SourceType:  "svn",
			TimerTrigger: &devops.TimerTrigger{
				Interval: "12345566",
			},
			SvnSource: &devops.SvnSource{
				Remote:       "https://api.svn.com/bcd",
				CredentialId: "svn",
				Excludes:     "truck",
				Includes:     "tag/*",
			},
		},
		{
			Name:        "",
			Description: "for test",
			ScriptPath:  "Jenkinsfile",
			SourceType:  "single_svn",
			TimerTrigger: &devops.TimerTrigger{
				Interval: "12345566",
			},
			SingleSvnSource: &devops.SingleSvnSource{
				Remote:       "https://api.svn.com/bcd",
				CredentialId: "svn",
			},
		},
	}

	for _, input := range inputs {
		outputString, err := createMultiBranchPipelineConfigXml("", input)
		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		output, err := parseMultiBranchPipelineConfigXml(outputString)

		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		if !reflect.DeepEqual(input, output) {
			t.Fatalf("input [%+v] output [%+v] should equal ", input, output)
		}
	}
}

func Test_MultiBranchPipelineCloneConfig(t *testing.T) {

	inputs := []*devops.MultiBranchPipeline{
		{
			Name:        "",
			Description: "for test",
			ScriptPath:  "Jenkinsfile",
			SourceType:  "git",
			GitSource: &devops.GitSource{
				Url:              "https://github.com/kubesphere/devops",
				CredentialId:     "git",
				DiscoverBranches: true,
				CloneOption: &devops.GitCloneOption{
					Shallow: false,
					Depth:   3,
					Timeout: 20,
				},
			},
		},
		{
			Name:        "",
			Description: "for test",
			ScriptPath:  "Jenkinsfile",
			SourceType:  "github",
			GitHubSource: &devops.GithubSource{
				Owner:                "kubesphere",
				Repo:                 "devops",
				CredentialId:         "github",
				ApiUri:               "https://api.github.com",
				DiscoverBranches:     1,
				DiscoverPRFromOrigin: 2,
				DiscoverPRFromForks: &devops.DiscoverPRFromForks{
					Strategy: 1,
					Trust:    1,
				},
				CloneOption: &devops.GitCloneOption{
					Shallow: false,
					Depth:   3,
					Timeout: 20,
				},
			},
		},
	}

	for _, input := range inputs {
		outputString, err := createMultiBranchPipelineConfigXml("", input)
		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		output, err := parseMultiBranchPipelineConfigXml(outputString)

		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		if !reflect.DeepEqual(input, output) {
			t.Fatalf("input [%+v] output [%+v] should equal ", input, output)
		}
	}

}

func Test_MultiBranchPipelineRegexFilter(t *testing.T) {

	inputs := []*devops.MultiBranchPipeline{
		{
			Name:        "",
			Description: "for test",
			ScriptPath:  "Jenkinsfile",
			SourceType:  "git",
			GitSource: &devops.GitSource{
				Url:              "https://github.com/kubesphere/devops",
				CredentialId:     "git",
				DiscoverBranches: true,
				RegexFilter:      ".*",
			},
		},
		{
			Name:        "",
			Description: "for test",
			ScriptPath:  "Jenkinsfile",
			SourceType:  "github",
			GitHubSource: &devops.GithubSource{
				Owner:                "kubesphere",
				Repo:                 "devops",
				CredentialId:         "github",
				ApiUri:               "https://api.github.com",
				DiscoverBranches:     1,
				DiscoverPRFromOrigin: 2,
				DiscoverPRFromForks: &devops.DiscoverPRFromForks{
					Strategy: 1,
					Trust:    1,
				},
				RegexFilter: ".*",
			},
		},
	}

	for _, input := range inputs {
		outputString, err := createMultiBranchPipelineConfigXml("", input)
		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		output, err := parseMultiBranchPipelineConfigXml(outputString)

		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		if !reflect.DeepEqual(input, output) {
			t.Fatalf("input [%+v] output [%+v] should equal ", input, output)
		}
	}

}

func Test_MultiBranchPipelineMultibranchTrigger(t *testing.T) {

	inputs := []*devops.MultiBranchPipeline{
		{
			Name:        "",
			Description: "for test",
			ScriptPath:  "Jenkinsfile",
			SourceType:  "github",
			GitHubSource: &devops.GithubSource{
				Owner:                "kubesphere",
				Repo:                 "devops",
				CredentialId:         "github",
				ApiUri:               "https://api.github.com",
				DiscoverBranches:     1,
				DiscoverPRFromOrigin: 2,
				DiscoverPRFromForks: &devops.DiscoverPRFromForks{
					Strategy: 1,
					Trust:    1,
				},
				RegexFilter: ".*",
			},
			MultiBranchJobTrigger: &devops.MultiBranchJobTrigger{
				CreateActionJobsToTrigger: "abc",
				DeleteActionJobsToTrigger: "ddd",
			},
		},
		{
			Name:        "",
			Description: "for test",
			ScriptPath:  "Jenkinsfile",
			SourceType:  "github",
			GitHubSource: &devops.GithubSource{
				Owner:                "kubesphere",
				Repo:                 "devops",
				CredentialId:         "github",
				ApiUri:               "https://api.github.com",
				DiscoverBranches:     1,
				DiscoverPRFromOrigin: 2,
				DiscoverPRFromForks: &devops.DiscoverPRFromForks{
					Strategy: 1,
					Trust:    1,
				},
				RegexFilter: ".*",
			},
			MultiBranchJobTrigger: &devops.MultiBranchJobTrigger{
				CreateActionJobsToTrigger: "abc",
			},
		},
		{
			Name:        "",
			Description: "for test",
			ScriptPath:  "Jenkinsfile",
			SourceType:  "github",
			GitHubSource: &devops.GithubSource{
				Owner:                "kubesphere",
				Repo:                 "devops",
				CredentialId:         "github",
				ApiUri:               "https://api.github.com",
				DiscoverBranches:     1,
				DiscoverPRFromOrigin: 2,
				DiscoverPRFromForks: &devops.DiscoverPRFromForks{
					Strategy: 1,
					Trust:    1,
				},
				RegexFilter: ".*",
			},
			MultiBranchJobTrigger: &devops.MultiBranchJobTrigger{
				DeleteActionJobsToTrigger: "ddd",
			},
		},
	}

	for _, input := range inputs {
		outputString, err := createMultiBranchPipelineConfigXml("", input)
		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		output, err := parseMultiBranchPipelineConfigXml(outputString)

		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}
		if !reflect.DeepEqual(input, output) {
			t.Fatalf("input [%+v] output [%+v] should equal ", input, output)
		}
	}

}
