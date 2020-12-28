package main

import (
	"fmt"
	"os"

	"github.com/sclevine/agouti"
)

// CloudLog の操作
type CloudLog struct {
	page   *agouti.Page
	driver *agouti.WebDriver
}

// NewCloudLog CloudLogのビルダーメソッド
func NewCloudLog() (*CloudLog, error) {
	options := agouti.ChromeOptions(
		"args", []string{
			"--headless",
			// "--disable-gpu", // 暫定的に必要らしいです。
			"--lang=ja",
		})
	c := new(CloudLog)
	c.driver /*const*/ = agouti.ChromeDriver(options)

	if err := c.driver.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return nil, err
	}

	var err error
	c.page, err = c.driver.NewPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return nil, err
	}

	return c, nil
}

// Close デストラクタ
func (c *CloudLog) Close() {
	c.driver.Stop()
}

// Login ログイン処理
func (c *CloudLog) Login() {

	fmt.Println("login page")
	c.page.Navigate("https://app.innopm.com/login.cgi")
	// fmt.Println(os.Getenv("CLOUDLOG_EMAIL"))
	// fmt.Println(os.Getenv("CLOUDLOG_PASSWORD"))
	c.page.FindByName("email").Fill(os.Getenv("CLOUDLOG_EMAIL"))
	c.page.FindByName("passwd").Fill(os.Getenv("CLOUDLOG_PASSWORD"))
	c.page.FindByButton("ログイン").Click()
	// time.Sleep(5000 * time.Millisecond)
}

func clockIn() {

}

func clockOut() {

}
func main() {
	cloudLog, _ := NewCloudLog()
	defer cloudLog.Close()

	cloudLog.Login()
}
