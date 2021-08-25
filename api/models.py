from django.db import models

# Create your models here.

class serve_url( models.Model ):
  user       = models.TextField(blank=False)
  url        = models.URLField(blank=True)
  created_at = models.DateTimeField(auto_now_add=True)

