from singleton_pattern import ServerConfigSingleton


# 方式一：使用模块实现单例，多线程安全
server_config = ServerConfigSingleton


if __name__ == "__main__":

    import threading


    def task(arg):
        obj = ServerConfigSingleton
        print(id(obj))


    for i in range(10):
        t = threading.Thread(target=task, args=[i, ])
        t.start()
