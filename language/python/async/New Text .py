import os
import urllib.request
from concurrent.futures import ThreadPoolExecutor,ProcessPoolExecutor
from concurrent.futures import as_completed
'''
def square(a):
    square = a
    no = yield square**2

def client(to):
    for x in range(to+1):
        yield from square(x)

for x in list(client(1000)):
    print(x)

def downloader(url):
    """
    Downloads the specified URL and saves it to disk
    """
    req = urllib.request.urlopen(url)
    filename = os.path.basename(url)
    ext = os.path.splitext(url)[1]
    if not ext:
        raise RuntimeError('URL does not contain an extension')
    with open(filename, 'wb') as file_handle:
        while True:
            chunk = req.read(1024)
            if not chunk:
                break
            file_handle.write(chunk)
    msg = 'Finished downloading {filename}'.format(filename=filename)
    return msg
def main(urls):
    """
    Create a thread pool and download specified urls
    """
    with ThreadPoolExecutor(max_workers=5) as executor:
        future = executor.map(downloader,urls)
        print(list(future))

if __name__ == '__main__':
    urls = ["http://www.irs.gov/pub/irs-pdf/f1040.pdf",
            "http://www.irs.gov/pub/irs-pdf/f1040a.pdf",
            "http://www.irs.gov/pub/irs-pdf/f1040ez.pdf",
            "http://www.irs.gov/pub/irs-pdf/f1040es.pdf",
            "http://www.irs.gov/pub/irs-pdf/f1040sb.pdf"]
    #main(urls)
'''
'''
from collections import deque
def check(no):
    if no%2 == 0:
        return no
    else: 
        pass

def worker(no):
    print(f'------------worker-{no}-started---------')
    q = deque([x for x in filter(check,range(20*no))])
    print(f'------------worker-{no}--End------------')
    return q

def Threading(tasks):
    max_workers = len(tasks)
    with ThreadPoolExecutor(max_workers) as executor:
        #futures = executor.map(worker,tasks)
        futures = [executor.submit(worker,task) for task in tasks]
        print(futures)
        for future in as_completed(futures):
            print(future,future.result())
    return (list(futures))
tasks = [x for x in range(1,2)]
print(Threading(tasks))

import asyncio
@asyncio.coroutine
def countdown(number, n):
    while n > 0:
        print('T-minus', n, '({})'.format(number))
        yield from asyncio.sleep(1)
        n -= 1

loop = asyncio.get_event_loop()
tasks = [
    asyncio.ensure_future(countdown("A", 2)),
    asyncio.ensure_future(countdown("B", 3))]
loop.run_until_complete(asyncio.wait(tasks))
loop.close()
help(asyncio.ensure_future)'''
''' Description of asyncio.Future with coroutine's yield from <FUTURE>.

future: That will happen so we schedule it , when that will done we can do further processing.
Coroutine: A function that can be suspend and resume if exp after Yield  from takes some time
           and let the event loop to do further work .

        So What if i say yield from <FUTURE>??? 
if <FUTURE> may or may not complete yield will suspend or resume without blocking.
And Event Loop run fine.
/// A Book's statement.
Using yield from with a future automatically takes care of waiting for it to finish,
without blocking the event loop — because in asyncio, yield from is used to give
control back to the event loop.
we can simply put whatever processing you would do after the future is done in the lines
that follow yield from my_future in your coroutine. That’s the big advantage of having
co‐routines: functions that can be suspended and resumed.

You don’t need my_future.result() because the the value of yield from expres‐
sion on a future is the result, e.g. result = yield from my_future.

We’ll now consider how yield from and the asyncioAPI brings together futures, tasks
and coroutines.
573
'''
import asyncio

@asyncio.coroutine
def sum():
    return (20+20)

@asyncio.coroutine
def coroute():
    a = yield from sum()

def sync_code():
    loop = asyncio.get_event_loop() 
    tasks = [coroute(),coroute()]
    done= asyncio.wait(tasks)#wrap coroutine in asyncio.task(future)
    help(loop.create_task)
    result = loop.run_until_complete(done)
    #help(loop.run_until_complete)
    '''wrap coroutine in asyncio.Task objects automatically, using asyncio.async internally
    The loop.run_until_complete function accepts a future or a coroutine. If it gets a
    coroutine, run_until_complete wraps it into a Task, similar to what {wait} does.
    '''
    loop.close()

    return list(result)
print(sync_code())

@asyncio.coroutine
def inner():
   return 'say'
   
#print(iscoroutinefunction(inner))
def outer():
    yield from inner()

def client():
    loop = asyncio.get_event_loop()
    #tasks = asyncio.wait([outer(),outer(),outer()])#list of Tasks (Futures) 
    tasks = asyncio.ensure_future(outer())
    result = loop.run_until_complete(tasks)
    print(result)
    loop.close()
    return result

#print(client())
'''									///// Most important /////
@asyncio.coroutine
def say(what, when):
    yield from asyncio.sleep(when)
    print(what)

@asyncio.coroutine
def main(loop,msg):
    print(msg)
    task1 = loop.create_task(say('first hello', 4))
    task2 = loop.create_task(say('second hello', 1))
    yield from asyncio.gather(task1,task2,loop=loop)
    print('close')
@asyncio.coroutine
def main1(loop,msg):
    print(msg)
    task1 = loop.create_task(say('third hello', 2))
    task2 = loop.create_task(say('fourth hello', 1))
    yield from asyncio.gather(task1,task2,loop=loop) #to note here it will include the tasks in main queue and run async-sly
    print('close')

loop = asyncio.get_event_loop()
task = [asyncio.ensure_future(main(loop,'first')),asyncio.ensure_future(main1(loop,'second'))]
loop.run_until_complete(asyncio.wait(task))
loop.close()


'''
