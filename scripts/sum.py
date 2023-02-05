import datetime, requests, time
end = '2023-02-05'
host = 'https://stats.permaswap.network'

def get_stats(host, date):
    url = '%s/stats?date=%s'%(host, date)
    return requests.get(url).json()

stats = {}
start = datetime.date(2022, 12, 12)
total = 0
i = 0
while True:
    date = start.strftime('%Y-%m-%d')
    if date == end:
        break
    
    data = get_stats(host, date)
    for u, v in data['user'].items():
        stats[u] = stats.get(u, 0) + v
        total += v
    
    i += 1
    print(i, date)
    start = start + datetime.timedelta(days=1)
    time.sleep(0.25)
    
stats = dict(sorted(stats.items(), key=lambda item: item[1], reverse=True))

print('total', total, 'i', i)
for u, v in stats.items():
    print(u, v)