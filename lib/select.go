package lib

import "github.com/AlecAivazis/survey/v2"

func SelectTarget(platforms []string) ([]string, error) {
	targets := make([]string, 0)
	prompt := &survey.MultiSelect{
		Message: "Please Select Target Platforms:",
		Options: platforms,
	}

	if err := survey.AskOne(prompt, &targets); err != nil && err.Error() != "interrupt" {
		return nil, err
	}

	return targets, nil
}
