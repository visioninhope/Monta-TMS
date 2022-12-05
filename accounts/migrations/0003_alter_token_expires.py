# Generated by Django 4.1.3 on 2022-12-05 02:30

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("accounts", "0002_token"),
    ]

    operations = [
        migrations.AlterField(
            model_name="token",
            name="expires",
            field=models.DateTimeField(
                blank=True,
                help_text="The date and time the token expires",
                null=True,
                verbose_name="Expires",
            ),
        ),
    ]
