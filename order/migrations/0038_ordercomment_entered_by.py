# Generated by Django 4.1.3 on 2022-11-21 20:54

import django.db.models.deletion
from django.conf import settings
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        migrations.swappable_dependency(settings.AUTH_USER_MODEL),
        ("order", "0037_alter_stop_sequence"),
    ]

    operations = [
        migrations.AddField(
            model_name="ordercomment",
            name="entered_by",
            field=models.ForeignKey(
                default=1,
                help_text="Entered By",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="order_comments",
                related_query_name="order_comment",
                to=settings.AUTH_USER_MODEL,
                verbose_name="Entered By",
            ),
            preserve_default=False,
        ),
    ]