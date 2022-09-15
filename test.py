from http import server
import requests
import os

if __name__ == '__main__':
    os.system('start ./server/main.exe')
    
    server_port = 8081
    clients = {}

    print('0. Выоход\n1. Добавить клиент\n2.Список клиентов\n3.Запросить ключ и установить клиенту\n')

    while ((choose := int(input('Чо делать: '))) != 0):
        if choose == 0:
            break
        if choose == 1:
            os.system(f"start ./client/main.exe {int(port:=input('Порт: '))}")
            clients[port] = {'key': ''}
        if choose == 2:
            print('\n'.join(list(map(str, clients.items()))))
        if choose == 3:
            try:
                json = requests.post(f"http://localhost:{server_port}/create-api-key").json()
                print('JSON:', json)
                resp = requests.post(f"http://localhost:{int(port:=input('Порт: '))}/set-api-key", json={'uuid': json['uuid']})
                if resp.status_code == 200:
                    clients[port]['key'] = json['uuid']
            except Exception as e:
                print('Error: ', e)
