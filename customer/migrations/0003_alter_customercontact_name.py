# Generated by Django 4.1.4 on 2022-12-16 02:40

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("customer", "0002_initial"),
    ]

    operations = [
        migrations.AlterField(
            model_name="customercontact",
            name="name",
            field=models.CharField(
                help_text="Contact name", max_length=150, verbose_name="Name"
            ),
        ),
    ]