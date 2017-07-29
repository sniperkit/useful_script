1.文档字符串可以通过name.__doc__来访问，其中name可以是包名也可以是函数名。
2.使用help函数可以获取函数或者包的文档字符串.
	help(os)
	hlep(os.open)
3.使用class的__bases__属性，可以获取类的基类.
4.使用isinstance方法来检查一个对象是否是一个类的实例.
5.想要知道某个对象属于某一个类，可以使用它的__class__属性.
6.使用hasattr方法可已检查某一个对象是否包含某一个属性。
7.使用hasattr(x.'__call__')来检查某个对象是否可以被调用（函数）.
8.使用setattr来为一个对象设置属性.
9.导入一个包以后，可以通过dir(package_name)来获取包内的所有常量，函数，类定义...
10. __init__方法是提供了一个类的构造函数.
11. __del__方法提供了析构函数的功能，但是其调用时机是随机的，因此不建议使用。
12.当前类和对象可以作为super函数的参数使用，调用函数返回的对象的任何方法都是调用超类的方法，而不是当前类的方法。
13.静态方法没又self参数，并且可以被类本身直接调用.
14.类成员方法在定义时需要名为cls的类似self的参数，类成员方法可以直接使用类名调用。但是cls参数时直接被绑定到类的.
15.想要告知程序本身是做为程序运行还是作为库导入到其他程序，可以使用__name__变量：在main程序中，__name__变量的值为__main__。而在导入模块中该值被设置为模块的名称。
16.获取python模块搜索路径的方法：(也就是说放置在这些路径中模块可以被python解释器找到)
	>>> import sys, pprint
	>>> pprint.pprint(sys.path)
17.可通过修改PYTHONPATH环境变量来告诉python解释器去那里找模块定义.(自定以模块路径)
18.包含模块代码的文件的名字要和模块名一样，再加上.py扩展名.
19.当模块存储在文件中时，包就时模块所在的目录。为了让python解释器将其作为包对待，该目录下必须包含一个__init__.py的文件。为了将模块放在包内，直接将模块放在包目录内即可.
20.查看模块包含的内容可以使用dir函数.它会将对象的所有特性（一起莫款的所有函数，类，变量等）列出。
21.__all__定义了一个模块的共有接口，它告诉python解释器：从模块导入所有名字代表什么含义.如果没有设置__all___ ， import * 语句会导入模块中所有不以__开头的全局名称.
22.模块的__file__属性可以用来获取模块的源码路径.
23. python3.6 -m pip install Pillow
24. sudo vim /etc/sudoers
	secure_path = /sbin:/bin:/usr/sbin:/usr/bin:/usr/local/bin