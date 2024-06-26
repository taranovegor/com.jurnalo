systemctl stop jurnalo
systemctl disable jurnalo
rm /etc/systemd/system/jurnalo.service
systemctl daemon-reload
systemctl reset-failed
