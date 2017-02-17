from distutils.core import setup

__version__ = '0.1'

setup(name='poohbear',
      version=__version__,
      description='Interface to poohbear server',
      author='Arthur Elliott',
      author_email='clownpriest@gmail.com',
      license="MIT",
      url='https://github.com/clownpriest/poohbear/lib/python',
      packages=["poohbear"],
      install_requires=['grpc', 'grpcio']
)
