/*
 * COPYRIGHT(c) 2024 Trenova
 *
 * This file is part of Trenova.
 *
 * The Trenova software is licensed under the Business Source License 1.1. You are granted the right
 * to copy, modify, and redistribute the software, but only for non-production use or with a total
 * of less than three server instances. Starting from the Change Date (November 16, 2026), the
 * software will be made available under version 2 or later of the GNU General Public License.
 * If you use the software in violation of this license, your rights under the license will be
 * terminated automatically. The software is provided "as is," and the Licensor disclaims all
 * warranties and conditions. If you use this license's text or the "Business Source License" name
 * and trademark, you must comply with the Licensor's covenants, which include specifying the
 * Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use
 * Grant, and not modifying the license in any other way.
 */

import {
  FeasibilityOperatorChoiceProps,
  ServiceIncidentControlChoiceProps,
  SeverityChoiceProps,
} from "@/lib/choices";
import { StatusChoiceProps } from "@/types";
import {
  CommentTypeFormValues,
  DelayCodeFormValues,
  DispatchControlFormValues,
  FeasibilityToolControlFormValues,
  FleetCodeFormValues,
  RateBillingTableFormValues,
} from "@/types/dispatch";
import { boolean, number, object, ObjectSchema, string } from "yup";
import { validateDecimal } from "../utils";

export const dispatchControlSchema: ObjectSchema<DispatchControlFormValues> =
  object().shape({
    recordServiceIncident: string<ServiceIncidentControlChoiceProps>().required(
      "Record Service Incident is required",
    ),
    gracePeriod: number().required("Grace Period is required"),
    deadheadTarget: number().required("Deadhead Target is required"),
    enforceWorkerAssign: boolean().required("Worker Assign is required"),
    trailerContinuity: boolean().required("Trailer Continuity is required"),
    dupeTrailerCheck: boolean().required("Dupe Trailer Check is required"),
    maintenanceCompliance: boolean().required(
      "Maintenance Compliance is required",
    ),
    maxShipmentWeightLimit: number().required(
      "Max Load Weight Limit is required",
    ),
    regulatoryCheck: boolean().required("Regulatory Check is required"),
    prevShipmentsOnHold: boolean().required(
      "Previous Shipments on Hold is required",
    ),
    workerTimeAwayRestriction: boolean().required(
      "Worker Time Away Restriction is required",
    ),
    tractorWorkerFleetConstraint: boolean().required(
      "Tractor Worker Fleet Constraint is required",
    ),
  });

export const feasibilityControlSchema: ObjectSchema<FeasibilityToolControlFormValues> =
  object().shape({
    mpwOperator: string<FeasibilityOperatorChoiceProps>().required(
      "Miles per week operator is required.",
    ),
    mpwCriteria: number()
      .moreThan(0, "Miles per week criteria must be greater than 0.")
      .required("Miles per week criteria is required."),
    mpdOperator: string<FeasibilityOperatorChoiceProps>().required(
      "Miles per day operator is required.",
    ),
    mpdCriteria: number()
      .moreThan(0, "Miles per day criteria must be greater than 0.")
      .required("Miles per day criteria is required."),
    mpgOperator: string<FeasibilityOperatorChoiceProps>().required(
      "Miles per gallon operator is required.",
    ),
    mpgCriteria: number()
      .moreThan(0, "Miles per gallon criteria must be greater than 0.")
      .required("Miles per gallon criteria is required."),
    otpOperator: string<FeasibilityOperatorChoiceProps>().required(
      "On-time performance operator is required.",
    ),
    otpCriteria: number()
      .moreThan(0, "On-time performance criteria must be greater than 0.")
      .required("On-time performance criteria is required."),
  });

export const delayCodeSchema: ObjectSchema<DelayCodeFormValues> =
  object().shape({
    status: string<StatusChoiceProps>().required("Status is required"),
    code: string()
      .required("Name is required")
      .max(4, "Code cannot be more than 4 characters"),
    description: string()
      .required("Description is required")
      .max(100, "Description cannot be more than 100 characters"),
    fCarrierOrDriver: boolean().required("Carrier or Driver is required"),
  });

export const fleetCodeSchema: ObjectSchema<FleetCodeFormValues> =
  object().shape({
    status: string<StatusChoiceProps>().required("Status is required"),
    code: string()
      .required("Name is required")
      .max(10, "Code cannot be more than 10 characters"),
    revenueGoal: string().test(
      "is-decimal",
      "Revenue Goal must be a decimal with no more than two decimal places",
      (value) => {
        if (value === undefined || value === null || value === "") {
          return true; // Passes validation for null, undefined, or empty string
        }
        return validateDecimal(value, 2);
      },
    ),
    deadheadGoal: string().test(
      "is-decimal",
      "Deadhead Goal must be a decimal with no more than two decimal places",
      (value) => {
        if (value === undefined || value === null || value === "") {
          return true; // Passes validation for null, undefined, or empty string
        }
        return validateDecimal(value, 2);
      },
    ),
    mileageGoal: string().test(
      "is-decimal",
      "Mileage Goal must be a decimal with no more than two decimal places",
      (value) => {
        if (value === undefined || value === null || value === "") {
          return true; // Passes validation for null, undefined, or empty string
        }
        return validateDecimal(value, 2);
      },
    ),
    description: string()
      .required("Description is required")
      .max(100, "Description cannot be more than 100 characters"),
    manager: string(),
  });

export const commentTypeSchema: ObjectSchema<CommentTypeFormValues> =
  object().shape({
    status: string<StatusChoiceProps>().required("Status is required"),
    name: string()
      .max(10, "Name cannot be more than 10 characters")
      .required("Name is required"),
    description: string()
      .max(100, "Description cannot be more than 100 characters")
      .required("Description is required"),
    severity: string<SeverityChoiceProps>().required("Severity is required"),
  });

export const rateBillingTableSchema: ObjectSchema<RateBillingTableFormValues> =
  object().shape({
    accessorialCharge: string().required("Accessorial Charge is required"),
    description: string().max(
      100,
      "Description cannot be more than 100 characters long",
    ),
    unit: number().required("Unit is required"),
    chargeAmount: number().required("Charge Amount is required"),
    subTotal: number().required("Subtotal is required"),
  });
