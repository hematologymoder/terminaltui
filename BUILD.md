# Building and Deploying Portfolio TUI

This guide shows you how to build and deploy your cute portfolio TUI so people can SSH into it! 💕

## Prerequisites

- Go 1.19+ installed
- A Linux server with SSH access
- Domain name (optional but recommended)

## Building

### 1. Local Development
```bash
# Clone/navigate to your portfolio directory
cd /path/to/portfoliotui

# Build locally for testing
make build

# Run locally
make run
```

### 2. Build for Linux Deployment
```bash
# Build statically compiled binary for Linux
make build-linux

# This creates a portable binary that works on most Linux systems
```

## Deployment Setup

### 1. Server Preparation

```bash
# On your server, create a dedicated user
sudo adduser portfolio --disabled-password --gecos ""

# Create portfolio directory
sudo mkdir -p /opt/portfolio
sudo chown portfolio:portfolio /opt/portfolio
```

### 2. Upload Files

```bash
# From your local machine, copy files to server
scp portfolio-tui your-server:/opt/portfolio/
scp config.example.json your-server:/opt/portfolio/

# SSH into server
ssh your-server

# Make binary executable
sudo chmod +x /opt/portfolio/portfolio-tui
sudo chown portfolio:portfolio /opt/portfolio/portfolio-tui
```

### 3. Configure Last.fm (Optional)

```bash
# On the server, set up Last.fm integration
sudo -u portfolio mkdir -p /home/portfolio/.config/portfolio-tui/
sudo -u portfolio cp /opt/portfolio/config.example.json /home/portfolio/.config/portfolio-tui/config.json

# Edit with your API credentials
sudo -u portfolio nano /home/portfolio/.config/portfolio-tui/config.json
```

```json
{
  "lastfm": {
    "api_key": "your_lastfm_api_key_here",
    "username": "your_lastfm_username"
  }
}
```

### 4. SSH Configuration

Create a custom SSH configuration to launch the portfolio:

```bash
# Edit SSH daemon config
sudo nano /etc/ssh/sshd_config
```

Add these lines:
```
# Portfolio user configuration
Match User portfolio
    ForceCommand /opt/portfolio/portfolio-tui
    AllowTcpForwarding no
    AllowAgentForwarding no
    PermitTunnel no
    X11Forwarding no
```

Restart SSH:
```bash
sudo systemctl restart sshd
```

### 5. Set Up SSH Keys

```bash
# Switch to portfolio user
sudo -u portfolio -i

# Create SSH directory
mkdir -p ~/.ssh
chmod 700 ~/.ssh

# Add your public key (or users who should have access)
echo "your-public-key-here" >> ~/.ssh/authorized_keys
chmod 600 ~/.ssh/authorized_keys
```

## DNS Setup (Optional)

If you have a domain, create a subdomain:

```bash
# Add DNS A record:
# portfolio.yourdomain.com -> your-server-ip
```

## Testing

Test the deployment:

```bash
# From any machine, SSH to your portfolio
ssh portfolio@your-server
# or with custom domain:
ssh portfolio@portfolio.yourdomain.com
```

You should see the cute portfolio TUI launch immediately! ✨

## Advanced: Systemd Service (Optional)

For monitoring and automatic restarts, create a systemd service:

```bash
# Create service file
sudo nano /etc/systemd/system/portfolio-tui.service
```

```ini
[Unit]
Description=Portfolio TUI SSH Service
After=network.target

[Service]
Type=forking
User=portfolio
Group=portfolio
WorkingDirectory=/opt/portfolio
ExecStart=/opt/portfolio/portfolio-tui
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
```

```bash
# Enable and start service
sudo systemctl enable portfolio-tui
sudo systemctl start portfolio-tui
```

## Security Notes

- The portfolio user is restricted and can only run the TUI
- No shell access is provided to portfolio users
- Consider using fail2ban to prevent SSH brute force attacks
- Keep your server updated regularly

## Customization

To update your portfolio:

1. Make changes locally
2. Build: `make build-linux`
3. Upload new binary: `scp portfolio-tui your-server:/opt/portfolio/`
4. Restart if using systemd: `sudo systemctl restart portfolio-tui`

## Troubleshooting

### Connection Issues
```bash
# Check SSH service
sudo systemctl status sshd

# Check portfolio user
sudo -u portfolio /opt/portfolio/portfolio-tui
```

### Last.fm Not Working
```bash
# Check config file
sudo -u portfolio cat /home/portfolio/.config/portfolio-tui/config.json

# Test API manually
curl "https://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=USERNAME&api_key=API_KEY&format=json&limit=1"
```

### Performance Issues
```bash
# Check system resources
htop

# Monitor SSH connections
sudo netstat -tnlp | grep :22
```

## Example Complete Setup Script

```bash
#!/bin/bash
# Quick deployment script

# Create user
sudo adduser portfolio --disabled-password --gecos ""

# Create directories
sudo mkdir -p /opt/portfolio
sudo chown portfolio:portfolio /opt/portfolio

# Copy binary (assuming you've built it)
sudo cp portfolio-tui /opt/portfolio/
sudo chmod +x /opt/portfolio/portfolio-tui
sudo chown portfolio:portfolio /opt/portfolio/portfolio-tui

# Configure SSH
echo "Match User portfolio" | sudo tee -a /etc/ssh/sshd_config
echo "    ForceCommand /opt/portfolio/portfolio-tui" | sudo tee -a /etc/ssh/sshd_config
echo "    AllowTcpForwarding no" | sudo tee -a /etc/ssh/sshd_config
echo "    AllowAgentForwarding no" | sudo tee -a /etc/ssh/sshd_config
echo "    PermitTunnel no" | sudo tee -a /etc/ssh/sshd_config
echo "    X11Forwarding no" | sudo tee -a /etc/ssh/sshd_config

# Restart SSH
sudo systemctl restart sshd

echo "Setup complete! Add your SSH public key to /home/portfolio/.ssh/authorized_keys"
```

## Have Fun! 🎀

Your cute portfolio TUI is now ready for the world! People can SSH in and see your amazing projects, skills, and currently playing music. ✨

For questions or issues, feel free to reach out through the contact section of your portfolio! 💕