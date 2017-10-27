package main

import (	
	"testing"
	/*"fmt"
	"os"
	"encoding/json"
	"net/http"
	"strings"	
	"io/ioutil"	
	"strconv"	*/
)

func Test1(t *testing.T){
	expected := "Udeain"
	actual := "udeain"
	if actual != expected {
	  t.Error("Test failed")
	}
}