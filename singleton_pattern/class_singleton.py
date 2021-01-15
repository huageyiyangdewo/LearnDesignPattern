import threading


class Test:
    def __init__(self):
        # 以下是多线程中模拟耗时操作
        import time
        time.sleep(1)
        pass


# 方式三：使用类，多线程下不安全
class Singleton(object):

    @staticmethod
    def get_instance():
        if not hasattr(Singleton, "_instance"):
            Singleton._instance = Test()
        return Singleton._instance


# 加锁保证线程安全
class SingletonLock(object):

    _instance_lock = threading.Lock()

    @staticmethod
    def get_instance():
        if not hasattr(Singleton, "_instance"):
            # 第一次实例化后，后面就不用加锁判断了，提高执行速度

            with SingletonLock._instance_lock:
                if not hasattr(Singleton, "_instance"):
                    Singleton._instance = Test()
                return Singleton._instance


if __name__ == "__main__":

    import threading
    import time

    def task(arg):
        obj = SingletonLock.get_instance()
        print(id(obj))


    for i in range(10):
        t = threading.Thread(target=task, args=[i, ])
        t.start()
