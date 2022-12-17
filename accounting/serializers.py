"""
COPYRIGHT 2022 MONTA

This file is part of Monta.

Monta is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Monta is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Monta.  If not, see <https://www.gnu.org/licenses/>.
"""

from typing import Any

from rest_framework import serializers
from rest_framework.exceptions import ValidationError

from accounting import models
from utils.serializers import GenericSerializer


class GeneralLedgerAccountSerializer(GenericSerializer):
    """
    General Ledger Account Serializer
    """

    is_active = serializers.BooleanField(default=True)
    account_type = serializers.ChoiceField(
        choices=models.GeneralLedgerAccount.AccountTypeChoices.choices
    )
    cash_flow_type = serializers.ChoiceField(
        choices=models.GeneralLedgerAccount.CashFlowTypeChoices.choices
    )
    account_sub_type = serializers.ChoiceField(
        choices=models.GeneralLedgerAccount.AccountSubTypeChoices.choices
    )
    account_classification = serializers.ChoiceField(
        choices=models.GeneralLedgerAccount.AccountClassificationChoices.choices
    )

    class Meta:
        """
        Metaclass for GeneralLedgerAccountSerializer
        """

        model = models.GeneralLedgerAccount
        fields = (
            "id",
            "is_active",
            "organization",
            "account_number",
            "description",
            "account_type",
            "cash_flow_type",
            "account_sub_type",
            "account_classification",
            "created",
            "modified",
        )
        read_only_fields = (
            "organization",
            "id",
            "code",
            "created",
            "modified",
        )


class RevenueCodeSerializer(serializers.ModelSerializer):
    """
    Revenue Code Serializer
    """

    expense_account = serializers.PrimaryKeyRelatedField(
        queryset=models.GeneralLedgerAccount.objects.filter(
            account_type=models.GeneralLedgerAccount.AccountTypeChoices.EXPENSE
        )
    )
    revenue_account = serializers.PrimaryKeyRelatedField(
        queryset=models.GeneralLedgerAccount.objects.filter(
            account_type=models.GeneralLedgerAccount.AccountTypeChoices.REVENUE
        )
    )

    class Meta:
        """
        Metaclass for RevenueCodeSerializer
        """

        model = models.RevenueCode
        fields = (
            "id",
            "organization",
            "code",
            "description",
            "expense_account",
            "revenue_account",
            "created",
            "modified",
        )

    def validate(self, data: Any) -> dict:
        """Validate the data

        Args:
            data (Any): The data to validate

        Returns:
            dict: The validated data
        """
        expense_account = data.get("expense_account")
        revenue_account = data.get("revenue_account")

        if (
            expense_account
            and expense_account.account_type
            != models.GeneralLedgerAccount.AccountTypeChoices.EXPENSE
        ):
            raise ValidationError(
                {"expense_account": "Entered account is not an expense account."}
            )
        if (
            revenue_account
            and revenue_account.account_type
            != models.GeneralLedgerAccount.AccountTypeChoices.REVENUE
        ):
            raise ValidationError(
                {"revenue_account": "Entered account is not a revenue account."}
            )

        return data