'''import asyncio
async def print_every_second():
    "Print seconds"
    while True:
        for i in range(60):
            print(i, 's')
            await asyncio.sleep(1)
async def print_every_minute():
    for i in range(1, 10):
        await asyncio.sleep(60)
        print(i, 'minute')
loop = asyncio.get_event_loop()
loop.run_until_complete(
    asyncio.gather(print_every_second(),
                   print_every_minute())
)
loop.close()
'''
import asyncio

@asyncio.coroutine
def say(what, when):
    yield from asyncio.sleep(when)
    print(what)

@asyncio.coroutine
def main(loop,msg):
    print(f'      opening {msg}')                               #--------result------
    task1 = loop.create_task(say('1st hello', 4))               #      opening first
    task2 = loop.create_task(say('2nd hello', 1))               #      opening second
    yield from asyncio.gather(task1,task2,loop=loop)            #2nd hello
    print(f'      closing {msg}')                               #4th hello
                                                                #3rd hello
                                                                #      closing second
                                                                #      closing first
@asyncio.coroutine
def main1(loop,msg):
                                                                #1st hello
    print(f'      opening {msg}')
    task1 = loop.create_task(say('3rd hello', 2))
    task2 = loop.create_task(say('4th hello', 1))
    yield from asyncio.gather(task1,task2,loop=loop)
    print(f'      closing {msg}')

loop = asyncio.get_event_loop()
task = [asyncio.ensure_future(main(loop,'first')),asyncio.ensure_future(main1(loop,'second'))]
loop.run_until_complete(asyncio.wait(task))
loop.close()