
import json
from datetime import datetime as dt
import requests
import psutil
from online.models import Online
from natural.models import Natural
from device.models import Devices
from history.models import History

 
def my_scheduled_job():

    #更新cpu
    cpu = psutil.cpu_percent()
    Natural(**{'name':'cpu', 'ava':cpu, 'total':'1'}).save()

    #更新online cpu
    for online in Online.objects.order_by("cpu"):
        try:
            body = requests.post(online.server + '/online', headers={"token":online.token})
            if body.status_code == 200:
                data = body.json()
                online.status = 2
                online.cpu    = '{}%'.format(int(data['result']['cpu']))
                online.memery = '{}%'.format(int(data['result']['memery']))
                online.disk   = '{}T'.format(round(data['result']['disk']['total']/(1024*1024*1024*1024),2))
                online.disk2  = '{}T'.format(round(data['result']['disk']['free']/(1024*1024*1024*1024),2))
            else:
                online.status = 3
        except Exception as e:
            print(e)
            online.status = 3
            online.update_time = dt.now()
        online.save()

        #获取发现设备
        try:
            body = requests.post(online.server + '/api/v1/network/monitor/devices', headers={"token":online.token}) 
            if body.status_code == 200:
                data = body.json()
                if data['status'] == 200:
                    for row in data['result']:
                        if not Devices.objects.filter(name = row['name']).filter(ip = row['ip']).first():
                            Devices(**row).save()
        except Exception as e:
            print(e)

        #获取流量记录
        try:
            body = requests.post(online.server + '/api/v1/network/monitor/history', headers={"token":online.token}) 
            if body.status_code == 200:
                data = body.json()
                print(data)
                if data['status'] == 200:
                    for row in data['result']:
                        del row['ID']
                        del row['CreatedAt']
                        History(**row).save()
        except Exception as e:
            print(e)

    