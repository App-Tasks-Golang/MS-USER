package adapters

import (
	"User-Service-Go/pkg/domain"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type AuthClient struct {
	baseURL string
}

func NewAuthClient(baseURL string) *AuthClient {
	return &AuthClient{baseURL: baseURL}
}

func (c *AuthClient) RegisterUser(UserRequest domain.UserRequest) (bool, string, error) {
	// Definir la URL del endpoint de autenticación
	url := fmt.Sprintf("%s/auth/register", c.baseURL)

	// Crear el cuerpo de la solicitud (request body) en formato JSON
	requestBody, err := json.Marshal(UserRequest)
	if err != nil {
		return false, "",fmt.Errorf("failed to marshal user data: %w", err)
	}

	// Crear un cliente HTTP con timeout
	client := &http.Client{Timeout: 5 * time.Second}

	// Crear la solicitud HTTP POST
	req, err := http.NewRequest("POST", url, strings.NewReader(string(requestBody)))
	if err != nil {
		return false, "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json") // Establecer el tipo de contenido

	// Enviar la solicitud
	resp, err := client.Do(req)
	if err != nil {
		return false, "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	// Procesar la respuesta
	if resp.StatusCode == http.StatusConflict {
		return false, string(body), nil
	}

	// Capturar otros códigos que no sean 201 (http.StatusCreated)
	if resp.StatusCode != http.StatusCreated {
		return false, string(body), nil
	}

	return true, "", nil // Usuario registrado con éxito
}


func (c *AuthClient) LoginUser(dataUser domain.UserRequest) (string, error) {
	// Definir la URL del endpoint de autenticación
	url := fmt.Sprintf("%s/auth/login", c.baseURL)

	// Crear el cuerpo de la solicitud (request body) en formato JSON
	requestBody, err := json.Marshal(dataUser)
	if err != nil {
		return "", fmt.Errorf("failed to marshal user data: %w", err)
	}

	// Crear un cliente HTTP con timeout
	client := &http.Client{Timeout: 5 * time.Second}

	// Crear la solicitud HTTP POST
	req, err := http.NewRequest("POST", url, strings.NewReader(string(requestBody)))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json") // Establecer el tipo de contenido

	// Enviar la solicitud
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

		// Verificar el código de estado HTTP
		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("login failed, status: %s", resp.Status)
		}
	
		// Leer la respuesta
		var response struct {
			Token string `json:"token"`
		}
	
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return "", fmt.Errorf("failed to decode response: %w", err)
		}
	
		// Devolver el token
		return response.Token, nil
}


func (c *AuthClient) LogoutUser(token string) error {
	// URL del microservicio de autenticación para cerrar sesión
	url := fmt.Sprintf("%s/auth/logout", c.baseURL)

	client := &http.Client{Timeout: 5 * time.Second}

	// Crear solicitud HTTP sin cuerpo JSON, solo el token en el header
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("error al crear la solicitud HTTP: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error al enviar la solicitud: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Println("Cuerpo de la respuesta:", string(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error al cerrar la sesión, código de estado: %d", resp.StatusCode)
	}

	return nil
}
