package util

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"pg-backend/models"
	"pg-backend/repository"
)

func SendPasswordResetEmail(user *models.User, token string) error {
	// get user info ->create sender -> parse template -> convert template to buffer and substitute values -> send email
	userInfo, err := repository.GetUserInfoByUserId(user.Id)
	if err != nil {
		return fmt.Errorf("failed to get user info: %v", err)
	}
	emailSender := NewEmailSender(nil)
	currentDir, _ := os.Getwd()
	templatePath := filepath.Join(currentDir, "template/email-templates/password_reset_link_template.html")
	template, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}
	var body bytes.Buffer
	data := struct {
		Username  string
		ResetLink string
	}{
		Username:  userInfo.Username,
		ResetLink: fmt.Sprintf("http://localhost:3000/reset-password?token=%s", token),
	}
	if err := template.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}
	err = emailSender.SendEmail(user.Email, "Reset Password", body.String())
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}
