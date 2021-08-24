from django.http import HttpResponse
from django.shortcuts import render

def index(request):
  return HttpResponse("It's home page, and <a href='/api'>api here</a>")

def home( request ):
  return render(request, 'index.html')