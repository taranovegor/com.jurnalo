cat <<EOF
   _                        _       _
  (_)                      | |     | |
   _ _   _ _ __ _ __   __ _| | ___ | |
  | | | | |  __|  _ \ / _  | |/ _ \\| |
  | | |_| | |  | | | | (_| | | (_) |_|
  | |\____|_|  |_| |_|\____|_|\___/(_)
 _/ |
|__/

jurnalo has been installed as a systemd service.

To start/stop jurnalo:

sudo systemctl start jurnalo
sudo systemctl stop jurnalo

To enable/disable jurnalo starting automatically on boot:

sudo systemctl enable jurnalo
sudo systemctl disable jurnalo

To reload jurnalo:

sudo systemctl restart jurnalo

To view jurnalo logs:

journalctl -f -u jurnalo

EOF
