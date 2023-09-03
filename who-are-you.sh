curl -s https://learn.reboot01.com/git/ralasfoor/bh-piscine.git | jq ' .[] | select( .id == 70) .name'
