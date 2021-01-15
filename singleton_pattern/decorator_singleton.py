import threading


# 方式二：使用装饰器实现单例，多线程不安全
def singleton(cls):
    _instance = {}

    def _singleton(*args, **kwargs):
        if cls not in _instance:
            _instance[cls] = cls(*args, **kwargs)
        return _instance[cls]

    return _singleton


# 加锁保证多线程下是安全的
def singleton_lock(cls):
    _instance = {}
    _instance_lock = threading.Lock()

    def _singleton(*args, **kwargs):
        with _instance_lock:
            if cls not in _instance:
                _instance[cls] = cls(*args, **kwargs)
        return _instance[cls]

    return _singleton


if __name__ == "__main__":

    import time

    @singleton_lock
    class A:

        def __init__(self, x):
            import time
            time.sleep(1)

            self.x = x


    def task(arg):
        obj = A(arg)
        print(id(obj))


    for i in range(10):
        t = threading.Thread(target=task, args=[i, ])
        t.start()
