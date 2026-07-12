# Auditor's Arsenal - SRE Command Reference

# NETWORK
ss -tanpi                    # -t TCP, -a all, -n numeric, -p process, -i internals | Connection debugging, latency, retransmits
ss -s                        # -s summary | Port exhaustion, TIME_WAIT accumulation
ethtool -S eth0 | grep -i drop  # -S stats, eth0 interface | NIC ring buffer overflow, hardware drops
tcpdump -i eth0 -nn -s0 -w /tmp/debug.pcap 'port 9092'  # -i iface, -nn no resolve, -s0 full pkt, -w file | Full capture for Wireshark
tcpdump -i eth0 'tcp[tcpflags] & (tcp-syn|tcp-rst|tcp-fin) != 0'  # tcpflags filter | Connection lifecycle (SYN/ACK/RST/FIN)
tcpdump -i eth0 -nn 'udp port 53'  # udp port 53 | DNS query debugging
nslookup <service>           # Quick DNS resolution test

# PROCESS & RESOURCES
pidstat -p <PID> 1           # -p PID, 1 sec interval | Per-process CPU, memory, I/O stats
pidstat -d -p <PID> 1        # -d disk I/O | Disk I/O per process, identify thrashing
lsof -p <PID> | grep REG     # -p PID, grep REG regular files | File descriptor leaks, open file count
ls -l /proc/<PID>/fd | wc -l # /proc/<PID>/fd file descriptors, wc -l count | Quick FD count check

# MEMORY & KERNEL
cat /proc/meminfo | grep -E 'MemAvailable|Buffers|Cached'  # grep -E extended regex | Real memory pressure, OOM investigation
cat /proc/vmstat | grep -E 'pgfault|pgmajfault'  # pgfault minor, pgmajfault major page faults | Memory pressure, disk-backed access
dmesg -T | grep -i 'oom\|kill'  # -T human timestamp, grep oom/kill | OOM killer activity, unexpected process deaths
cat /proc/sys/vm/dirty_ratio # Writeback pressure tuning

# DISK & FILESYSTEM
df -ih                       # -i inodes, -h human readable | Inode exhaustion, "no space" with free disk
du -sh /* 2>/dev/null | sort -hr | head -10  # -s summary, -h human, /* all root, sort -hr human reverse | Find disk space hogs
find /var/log -type f -size +100M -exec ls -lh {} \;  # -type f files, -size +100M >100MB | Find giant unrotated log files
iostat -x 1                  # -x extended, 1 interval | Disk utilization, queue depth, await time

# CONTAINER DEBUGGING
docker inspect <container> --format '{{.State.Pid}}'  # --format Go template for PID | Get host PID of container
PID=$(docker inspect -f '{{.State.Pid}}' <container>) && nsenter -t $PID -n ip addr show  # nsenter -t target, -n network ns | Container's network interfaces
PID=$(docker inspect -f '{{.State.Pid}}' <container>) && nsenter -t $PID -m cat /proc/mounts  # -m mount ns | Container's mount points
crictl ps -a                 # -a all containers | List containers via containerd (K8s)

# PERFORMANCE PROFILING
perf top -p <PID>            # -p PID | Real-time CPU hotspot identification
strace -p <PID> -c           # -p PID, -c summary | Syscall analysis, kernel time breakdown
cat /proc/<PID>/stack        # Kernel stack trace, D state debugging

# KILL-SHOT COMBOS

## Pod Slow
ss -tanpi | grep <port>
pidstat -p <PID> 1
ls -l /proc/<PID>/fd | wc -l
dmesg -T | tail -20

## Kafka Consumer Lag
ss -tanpi | grep 9092
pidstat -d -p <consumer_pid> 1
iostat -x 1
tcpdump -i eth0 -nn -s0 -w /tmp/kafka-lag.pcap 'port 9092'

## DNS Slow
nslookup <service>
tcpdump -i eth0 -nn 'udp port 53' -w /tmp/dns-debug.pcap
cat /etc/resolv.conf
resolvectl statistics