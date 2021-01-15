import threading


# 方式五：基于 metaclass 实现,多线程不安全
class Singleton(type):

    def __call__(cls, *args, **kwargs):
        if not hasattr(cls, "_instance"):
            cls._instance = super(Singleton, cls).__call__(*args, **kwargs)
        return cls._instance


# 加锁保证多线程安全
class SingletonLock(type):

    _instance_lock = threading.Lock()

    def __call__(cls, *args, **kwargs):
        if not hasattr(cls, "_instance"):
            with cls._instance_lock:
                if not hasattr(cls, "_instance"):
                    cls._instance = super(SingletonLock, cls).__call__(*args, **kwargs)
        return cls._instance


class T(metaclass=SingletonLock):
    # 指定创建T的type为SingletonType
    def __init__(self, name):
        import time
        time.sleep(1)
        self.name = name


if __name__ == "__main__":

    import threading
    import time

    def task(arg):
        obj = T(arg)
        print(id(obj))


    for i in range(10):
        t = threading.Thread(target=task, args=[i, ])
        t.start()
