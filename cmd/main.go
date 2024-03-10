package main

import (
  "github.com/adalrikus/go-templ/pkg/slick"
  "github.com/adalrikus/go-templ/pkg/views/profile"
)

func main() {
  s := slick.New()
  s.Get("/profile", HandleUserProfile)
  s.Start(":3000")
}

func HandleUserProfile(c *slick.Context) error {
  return c.Render(profile.Index())
}
