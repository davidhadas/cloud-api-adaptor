#!/bin/bash
#
# Copyright (c) 2024 IBM Corporation
#
# SPDX-License-Identifier: Apache-2.0

set -euo pipefail

TEST_DIR=$(cd "$(dirname "$(realpath "$0")")/../"; pwd)

VERSIONS_YAML_PATH=$(realpath "${TEST_DIR}/../versions.yaml")

TRUSTEE_OPERATOR_REPO=$(yq -e '.git.trustee-operator.url' "${VERSIONS_YAML_PATH}")
TRUSTEE_OPERATOR_VERSION=$(yq -e '.git.trustee-operator.reference' "${VERSIONS_YAML_PATH}")

echo "${TRUSTEE_OPERATOR_REPO}, ${TRUSTEE_OPERATOR_VERSION}"

rm -rf "${TEST_DIR}/trustee_operator"
git clone "${TRUSTEE_OPERATOR_REPO}" "${TEST_DIR}/trustee_operator"
pushd "${TEST_DIR}/trustee_operator"
git checkout "${TRUSTEE_OPERATOR_VERSION}"
TRUSTEE_OPERATOR_SHA="$(git rev-parse HEAD)"
popd
