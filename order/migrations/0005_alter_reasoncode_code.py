# Generated by Django 4.1.5 on 2023-01-13 21:33

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("order", "0004_reasoncode_is_active"),
    ]

    operations = [
        migrations.AlterField(
            model_name="reasoncode",
            name="code",
            field=models.CharField(
                help_text="Code of the Reason Code",
                max_length=5,
                unique=True,
                verbose_name="Code",
            ),
        ),
    ]
