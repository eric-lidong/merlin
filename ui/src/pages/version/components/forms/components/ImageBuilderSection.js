import { FormLabelWithToolTip, useOnChangeHandler } from "@caraml-dev/ui-lib";
import {
  EuiFieldText,
  EuiFlexGroup,
  EuiFlexItem,
  EuiFormRow,
  EuiTitle,
} from "@elastic/eui";


export const ImageBuilderSection = ({
  imageBuilderResourceConfig,
  onChangeHandler,
  errors = {},
}) => {
  const {onChange} = useOnChangeHandler(onChangeHandler);

  return (
    <>
      <EuiTitle size="xs">
        <h4>{"Image Building Resources"}</h4>
      </EuiTitle>
    <EuiFlexGroup direction="row">
      <EuiFlexItem>
        <EuiFormRow
          label={
            <FormLabelWithToolTip
              label="Image Builder CPU"
              content="To set a higher CPU above platform default for Image Building"
            />
          }
          isInvalid={!!errors.cpu_request}
          error={errors.cpu_request}
          fullWidth
        >
          <EuiFieldText
            placeholder="500m"
            value={imageBuilderResourceConfig?.cpu_request}
            onChange={(e) => onChange("cpu_request")(e.target.value)}
            isInvalid={!!errors.cpu_request}
            name="cpu"
            fullWidth
          />
        </EuiFormRow>
      </EuiFlexItem>

      <EuiFlexItem>
        <EuiFormRow
          label={
            <FormLabelWithToolTip
              label="Image Builder Memory"
              content="To set a higher Mem above platform default for Image Building"
            />
          }
          isInvalid={!!errors.memory_request}
          error={errors.memory_request}
          fullWidth
        >
          <EuiFieldText
            placeholder="500Mi"
            value={imageBuilderResourceConfig?.memory_request}
            onChange={(e) => onChange("memory_request")(e.target.value)}
            isInvalid={!!errors.memory_request}
            name="memory"
            fullWidth
          />
        </EuiFormRow>
      </EuiFlexItem>
    </EuiFlexGroup>
    </>
  )
}