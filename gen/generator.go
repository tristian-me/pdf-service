package gen

import (
	"os"
	"os/exec"

	"pdf-service/utils"
)

const tmpFolder = "./tmp"

// ConvertFromFile handles the conversion from HTML to PDF. If the converion was
// successfull, the new filename will be returned with a nil [error].
func ConvertFromFile(fileName string) (string, error) {
	var err error

	chromeExec, err := exec.LookPath("chromium")
	if err != nil {
		return "", err
	}

	newFileName := utils.RandomString(30) + ".pdf"

	args := []string{
		chromeExec,
		"--headless",
		"--disable-gpu",
		"--print-to-pdf-no-header"
		"--print-to-pdf=" + tmpFolder + "/" + newFileName,
		tmpFolder + "/" + fileName,
	}

	cmd := &exec.Cmd{
		Path:   chromeExec,
		Args:   args,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	return newFileName, nil
}
