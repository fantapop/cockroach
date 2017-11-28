import Select from "react-select";
import React from "react";
import _ from "lodash";

import "./dropdown.styl";

import {leftArrow, rightArrow} from "src/views/shared/components/icons";
import { trustIcon } from "src/util/trust";

export interface DropdownOption {
  value: string;
  label: string;
}

export enum ArrowDirection {
  LEFT, RIGHT,
}

interface DropdownOwnProps {
  title: string;
  selected: string;
  options: DropdownOption[];
  onChange?: (selected: DropdownOption) => void; // Callback when the value changes.
  // If onArrowClick exists, don't display the arrow next to the dropdown,
  // display left and right arrows to either side instead.
  onArrowClick?: (direction: ArrowDirection) => void;
  // Disable any arrows in the arrow direction array.
  disabledArrows?: ArrowDirection[];
}

/**
 * Dropdown component that uses the URL query string for state.
 */
export default class Dropdown extends React.Component<DropdownOwnProps, {}> {
  render() {
    const {selected, options, onChange, onArrowClick, disabledArrows} = this.props;
    let className = "dropdown";
    if (onArrowClick) {
      className += " dropdown--side-arrows";
    }

    let leftClassName = "dropdown__side-arrow";
    if (_.includes(disabledArrows, ArrowDirection.LEFT)) {
      leftClassName += " dropdown__side-arrow--disabled";
    }

    let rightClassName = "dropdown__side-arrow";
    if (_.includes(disabledArrows, ArrowDirection.RIGHT)) {
      rightClassName += " dropdown__side-arrow--disabled";
    }

    return <div className={className}>
      {/* TODO (maxlang): consider moving arrows outside the dropdown component */}
      <span
        className={leftClassName}
        dangerouslySetInnerHTML={trustIcon(leftArrow)}
        onClick={() => this.props.onArrowClick(ArrowDirection.LEFT)}>
      </span>
      <span className="dropdown__title">{this.props.title}{this.props.title ? ":" : ""}</span>
      <Select className="dropdown__select" clearable={false} searchable={false} options={options} value={selected} onChange={onChange} />
      <span
        className={rightClassName}
        dangerouslySetInnerHTML={trustIcon(rightArrow)}
        onClick={() => this.props.onArrowClick(ArrowDirection.RIGHT)}>
      </span>
    </div>;
  }
}
