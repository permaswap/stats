import datetime, requests
end = '2022-12-31'
host = 'https://stats.permaswap.network'

def get_stats(host, date):
    url = '%s/stats?date=%s'%(host, date)
    return requests.get(url).json()

stats = {}
start = datetime.date(2022, 12, 22)
while True:
    date = start.strftime('%Y-%m-%d')
    data = get_stats(host, date)
    for u, v in data['user'].items():
        stats[u] = stats.get(u, 0) + v
    
    if date == end:
        break
    print(date)
    start = start + datetime.timedelta(days=1)

stats = dict(sorted(stats.items(), key=lambda item: item[1], reverse=True))

for u, v in stats.items():
    print(u, v)

    