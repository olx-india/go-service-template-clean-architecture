package redis

// LimitAvailable returns the remaining limit for a user.
// Adopter: replace this placeholder with the rate-limit backend operation.
func (p *Provider) LimitAvailable(_ int) (int, error) {
	if p == nil || p.client == nil {
		return 0, errRedisNotConfigured
	}
	return 0, nil
}

// ResetLimit resets the limit for a user.
// Adopter: replace this placeholder with the rate-limit backend operation.
func (p *Provider) ResetLimit(_ int) error {
	if p == nil || p.client == nil {
		return errRedisNotConfigured
	}
	return nil
}
