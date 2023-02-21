import { useState } from 'react';

import './Cell.css';

type CellProps = {
    rowIdx: number;
    colIdx: number;
    isAlive: boolean;
    onClickKill: (rowIdx: number, colIdx: number) => void
    onClickSpawn: (rowIdx: number, colIdx: number) => void
  };

const Cell = ({ rowIdx, colIdx, isAlive, onClickKill, onClickSpawn }: CellProps) => {

    return (
        <div
        className={`cell ${isAlive ? "cellAliveColor" : "cellDeadColor"}`}
        onClick={e => isAlive ? onClickKill(rowIdx, colIdx) : onClickSpawn(rowIdx, colIdx)}/>
    )
}

export default Cell
