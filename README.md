

1. **Generate SSH Key:**
   Run the following command to generate a new SSH key pair:
   ```sh
   ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
   ```
   Replace `"your_email@example.com"` with your actual GitHub email address. Press Enter to accept the default file location (`~/.ssh/id_rsa`) and optionally set a passphrase for added security.

2. **Add SSH Key to SSH Agent:**
   Start the SSH agent in your Codespace:
   ```sh
   eval "$(ssh-agent -s)"
   ```
   Add your SSH private key to the SSH agent:
   ```sh
   ssh-add ~/.ssh/id_rsa
   ```

3. **Copy SSH Key to Clipboard:**
   Display the SSH public key (`id_rsa.pub`):
   ```sh
   cat ~/.ssh/id_rsa.pub
   ```
   Copy the entire output of this command.

4. **Add SSH Key to GitHub:**
   - Go to your GitHub account settings.
   - Navigate to **Settings > SSH and GPG keys**.
   - Click on **New SSH key** or **Add SSH key**.
   - Paste your SSH key into the "Key" field.
   - Give your key a descriptive title, like "Codespace SSH Key".
   - Click **Add SSH key**.

5. **Verify SSH Configuration:**
   Confirm that your `~/.ssh/config` file exists and is correctly configured:
   ```sh
   nano ~/.ssh/config
   ```
   Add the following configuration (if it doesn't already exist):
   ```
   Host github.com
     HostName github.com
     User git
     IdentityFile ~/.ssh/id_rsa
   ```
   Save the file (`Ctrl+O`, `Enter`, `Ctrl+X` in nano).

6. **Push Changes to GitHub:**
   Update your remote URL to use SSH and push your changes:
   ```sh
   git remote set-url origin git@github.com:Galadon123/WebSocket-basics.git
   git push -u origin main
   ```

These steps should resolve the SSH authentication issue and allow you to push changes from your GitHub Codespace to your repository. If you encounter any issues or need further assistance, feel free to ask!