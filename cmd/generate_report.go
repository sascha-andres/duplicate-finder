/*
Copyright © 2021 Sascha Andres <sascha.andres@outlook.com>

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
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"livingit.de/code/dupfinder/reporter"
	"livingit.de/code/dupfinder/scanner"
)

func generateReport(logger *logrus.Entry, worker *scanner.Scanner) {
	report, err := reporter.GetReporter(viper.GetString("reporter.type"))
	if err != nil {
		logger.Fatalf("error creating reporter: %s", err)
	}
	err = report.Setup()
	if err != nil {
		logger.Errorf("error setting up the reporter: %s", err)
	}
	err = worker.Report(report)
	if err != nil {
		logger.Fatalf("error writing report: %s", err)
	}
	err = report.Finish()
	if err != nil {
		logger.Errorf("error finalizing the report: %s", err)
	}
}
