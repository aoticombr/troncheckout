package main

import "time"

type TokenTenant struct {
	TenantId  string    `json:"tenantId"`
	Nome      string    `json:"nome"`
	UsuarioId string    `json:"usuarioId"`
	Modulo    int       `json:"modulo"`
	Endereco  string    `json:"endereco"`
	Ativo     bool      `json:"ativo"`
	Role      string    `json:"role"`
	CriadoEm  time.Time `json:"criadoEm"`
}

type TokenData struct {
	Token        string        `json:"token"`
	RefreshToken string        `json:"refreshToken"`
	Expires      int           `json:"expires"`
	Tenants      []TokenTenant `json:"tenants"`
}
