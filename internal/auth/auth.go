package auth

// Auth workflow:
// 1. Client requests auth page link
// 2. Server returns auth page link (for now it's GitHub OAuth application auth page)
// 3. Client opens auth page link in browser
// 4. User (logs in to GitHub and)? grants access to the application
// 5. GitHub redirects user to callback URL with code
// 6. Server exchanges code for access token
// 7. Server checks via GitHub API if user is a member of the organization defined in config
// 8. If true, server (creates user in database and)? returns JWT token

// JWT generation workflow:
// 1. Server takes secret key from Fly.io secret store (Fly.io provides them via environment variables)
// 2. Server signs JWT token with secret key
