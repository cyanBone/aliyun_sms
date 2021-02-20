package aliyun_sms

import (
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type Client struct {
	accessKeyId     string
	accessKeySecret string
	signName        string
	scheme          string
}

func NewClient(accessKeyId string, accessKeySecret string, signName string, scheme string) *Client {
	if scheme == "" {
		scheme = "https"
	}
	return &Client{accessKeyId: accessKeyId, accessKeySecret: accessKeySecret, signName: signName, scheme: scheme}
}

func (c *Client) Send(tel, template, content string) error {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", c.accessKeyId, c.accessKeySecret)
	if err != nil {
		return err
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = c.scheme
	request.PhoneNumbers = tel
	request.SignName = c.signName
	request.TemplateCode = template
	request.TemplateParam = content
	sms, err := client.SendSms(request)
	if err != nil {
		return err
	}

	if sms.Code != "OK" {
		return errors.New(sms.Message)
	}

	return nil
}
