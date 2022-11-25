# Generated by Django 4.1.3 on 2022-11-21 20:22

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("order", "0036_ordercomment"),
    ]

    operations = [
        migrations.AlterField(
            model_name="stop",
            name="sequence",
            field=models.PositiveIntegerField(
                default=1,
                help_text="The sequence of the stop in the movement.",
                verbose_name="Sequence",
            ),
        ),
    ]