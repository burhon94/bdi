package main

import "testing"

func TestFunc_No_Param(t *testing.T)  {
	err := callMe(sample)
	if err != nil {
		t.Fatalf("error just be nil: %v", err)
	}
}

func TestFunc_No_Param_Err(t *testing.T)  {
	err := callMe(5)
	if err == nil {
		t.Fatalf("just be error: %v", err)
	}
}

func TestFunc_In_Param(t *testing.T)  {
	err := callMeAdvanced(sample2, "Param_In")
	if err != nil {
		t.Fatalf("error just be nil: %v", err)
	}
}

func TestFunc_In_Param_Err(t *testing.T)  {
	err := callMeAdvanced(false)
	if err == nil {
		t.Fatalf("just be error: %v", err)
	}
}

func TestFunc_In_Out_Param(t *testing.T)  {
	param, err := callMeAdvancedWithResult(sample3, "param_out")
	if err != nil {
		t.Fatalf("error just be nil: %v", err)
	}

	if param != "param_out" {
		t.Fatalf("param just be \"param_out\": %s", param)
	}
}

func TestFunc_In_Out_Param_Err(t *testing.T)  {
	_, err := callMeAdvancedWithResult("sample3", "param_out")
	if err == nil {
		t.Fatalf("just have error: %v", err)
	}
}