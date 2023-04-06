# Generated by Django 4.1.7 on 2023-04-03 17:06

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("kubectl", "0001_initial"),
    ]

    operations = [
        migrations.AlterField(
            model_name="kubeconfiguration",
            name="host",
            field=models.URLField(
                default="http://localhost",
                help_text="The host URL to use when communicating with the API server.",
                max_length=255,
                verbose_name="Host",
            ),
        ),
    ]