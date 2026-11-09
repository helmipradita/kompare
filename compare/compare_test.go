package compare

import (
	"os"
	"testing"

	"kompare/cli"
	"kompare/connect"
	"kompare/mock"

	"k8s.io/client-go/tools/clientcmd"
)

func TestCompareResourcesTargetClientset(t *testing.T) {
	_, _, tempKubeconfig := mock.SetupTestEnvironment()
	defer tempKubeconfig.Close()

	tempKubeconfigByte, err := os.ReadFile(tempKubeconfig.Name())
	if err != nil {
		t.Fatalf("Error reading kubeconfig file: %v", err)
	}
	kubeconfig, err := clientcmd.Load(tempKubeconfigByte)
	if err != nil {
		t.Fatalf("Error loading kubeconfig: %v", err)
	}

	var testContext string
	for context := range kubeconfig.Contexts {
		testContext = context
		break
	}

	configPath := tempKubeconfig.Name()
	sourceClientset, err := connect.ConnectToSource(testContext, &configPath)
	if err != nil {
		t.Fatalf("Error connecting to source: %v", err)
	}
	targetClientset, err := connect.ContextSwitch(testContext, &configPath)
	if err != nil {
		t.Fatalf("Error connecting to target: %v", err)
	}

	args := cli.ArgumentsReceivedValidated{
		FiltersForObject: "",
	}

	t.Run("CompareCronJobs", func(t *testing.T) {
		diff, err := CompareCronJobs(sourceClientset, targetClientset, "namespace2", args)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(diff) == 0 {
			t.Fatalf("expected non-empty comparison results")
		}
		for _, d := range diff {
			if len(d.Diff) != 0 {
				t.Errorf("expected no diff on identical clusters, got %v", d.Diff)
			}
		}
	})

	t.Run("CompareHPAs", func(t *testing.T) {
		diff, err := CompareHPAs(sourceClientset, targetClientset, "namespace2", args)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(diff) == 0 {
			t.Fatalf("expected non-empty comparison results")
		}
		for _, d := range diff {
			if len(d.Diff) != 0 {
				t.Errorf("expected no diff on identical clusters, got %v", d.Diff)
			}
		}
	})

	t.Run("CompareIngresses", func(t *testing.T) {
		diff, err := CompareIngresses(sourceClientset, targetClientset, "namespace2", args)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(diff) == 0 {
			t.Fatalf("expected non-empty comparison results")
		}
		for _, d := range diff {
			if len(d.Diff) != 0 {
				t.Errorf("expected no diff on identical clusters, got %v", d.Diff)
			}
		}
	})

	t.Run("CompareServices", func(t *testing.T) {
		diff, err := CompareServices(sourceClientset, targetClientset, "namespace2", args)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(diff) == 0 {
			t.Fatalf("expected non-empty comparison results")
		}
		for _, d := range diff {
			if len(d.Diff) != 0 {
				t.Errorf("expected no diff on identical clusters, got %v", d.Diff)
			}
		}
	})

	t.Run("CompareSecrets", func(t *testing.T) {
		diff, err := CompareSecrets(sourceClientset, targetClientset, "namespace2", args)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(diff) == 0 {
			t.Fatalf("expected non-empty comparison results")
		}
		if diff[0].Name != "secret1" {
			t.Errorf("expected secret1 from ListSecrets, got %s", diff[0].Name)
		}
		for _, d := range diff {
			if len(d.Diff) != 0 {
				t.Errorf("expected no diff on identical clusters, got %v", d.Diff)
			}
		}
	})

	t.Run("CompareServiceAccounts", func(t *testing.T) {
		diff, err := CompareServiceAccounts(sourceClientset, targetClientset, "namespace2", args)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(diff) == 0 {
			t.Fatalf("expected non-empty comparison results")
		}
		if diff[0].Name != "serviceaccount1" {
			t.Errorf("expected serviceaccount1 from ListServiceAccounts, got %s", diff[0].Name)
		}
		for _, d := range diff {
			if len(d.Diff) != 0 {
				t.Errorf("expected no diff on identical clusters, got %v", d.Diff)
			}
		}
	})
}
