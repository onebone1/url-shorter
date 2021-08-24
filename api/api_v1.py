from django.http import HttpResponse, JsonResponse

def version( request ):
  print( request.headers )
  version_info = {}
  version_info['version'] = '1.02'
  return JsonResponse( version_info )


def _URL( request ):
  return HttpResponse("Test")