/**
 * Copyright 2020 The Merlin Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { EuiFormRow, EuiSuperSelect } from "@elastic/eui";
import PropTypes from "prop-types";
import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { useMerlinApi } from "../../../../hooks/useMerlinApi";

export const ServiceAccountSelect = ({ selected, onChange }) => {
  const { projectId } = useParams();
  const [options, setOptions] = useState([]);

  const [{ data: secrets }] = useMerlinApi(
    `/projects/${projectId}/secrets`,
    {},
    [],
  );

  useEffect(() => {
    if (secrets) {
      const options = [];
      secrets
        .sort((a, b) => (a.name > b.name ? -1 : 1))
        .forEach((secret) => {
          options.push({
            value: secret.name,
            inputDisplay: secret.name,
            dropdownDisplay: secret.name,
          });
        });
      setOptions(options);
    }
  }, [secrets]);

  return (
    <EuiFormRow
      label="Select Service Account *"
      helpText={
        <p>
          Choose service account that have access to data source and sink. You
          can add new service account from Project's{" "}
          <a
            href={`/merlin/projects/${projectId}/settings/secrets-management`}
            target="_blank"
            rel="noopener noreferrer"
          >
            Secrets Management
          </a>{" "}
          page.
        </p>
      }
    >
      <EuiSuperSelect
        options={options}
        valueOfSelected={selected ? selected : ""}
        onChange={onChange}
        hasDividers
      />
    </EuiFormRow>
  );
};

ServiceAccountSelect.propTypes = {
  selected: PropTypes.string,
  onChange: PropTypes.func,
};
