package domain

import (
	"fmt"
	"mercado-livre-integration/pkg/infrastructure/client"
	"time"
)

func refreshTokenTask(m client.MercadoLivre) {
	Segovinha := make(map[string]string)

	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for range ticker.C {
			refreshToken, ok := Segovinha["refreshToken"]
			if ok {
				refreshToken, err := m.CreateRefreshToken(refreshToken)
				if err != nil {
					Segovinha["refreshToken"] = refreshToken.RefreshToken
				}
			} else {
				fmt.Println("refresh token nao localizado")
			}

		}
	}()
	//ticker.Stop()
}
