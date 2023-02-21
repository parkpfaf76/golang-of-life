import { useState, useEffect } from 'react';

import { GetCurrentGameStateGrid, UpdateToNextFrame, KillCell, SpawnCell, ClearCurrentGameState } from "../wailsjs/go/models/GameState";

import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import PlayArrow from '@mui/icons-material/PlayArrow';
import Pause from '@mui/icons-material/Pause';
import RestartAlt from '@mui/icons-material/RestartAlt';

import Cell from './components/Cell';

import './App.css';

const sleep = (ms: number) => new Promise((r) => setTimeout(r, ms));

function App() {
    const [currStateArray, setCurrStateArray] = useState(Array<Array<boolean>>());
    const [isPlaying, setIsPlaying] = useState(false)
    const [tick, setTick] = useState(0)


    useEffect(() => {
        return () => {
            GetCurrentGameStateGrid().then(res => setCurrStateArray(res))
        }
    }, [])

    useEffect(() => {
        const interval = setInterval(() => setTick(tick + 1), 500);

        if (isPlaying) {
            UpdateToNextFrame()
            GetCurrentGameStateGrid().then(res => setCurrStateArray(res))
        }

        return () => {
            clearInterval(interval);
        }
    }, [tick])

    let clearCurrentGameState = () => {
        console.log("RESET CLICKED")
        ClearCurrentGameState()
        GetCurrentGameStateGrid().then(res => setCurrStateArray(res))
    }

    let onClickKill = (rowIdx: number, colIdx: number) => {
        KillCell(rowIdx, colIdx)
        GetCurrentGameStateGrid().then(res => setCurrStateArray(res))
    }

    let onClickSpawn = (rowIdx: number, colIdx: number) => {
        SpawnCell(rowIdx, colIdx)
        GetCurrentGameStateGrid().then(res => setCurrStateArray(res))
    }

    let gameStateGridBody = currStateArray.map((row, i) =>
        <div className='gridRow'> {row.map((isCellAlive: boolean, j: number) => <Cell rowIdx={i} colIdx={j} isAlive={isCellAlive} onClickKill={onClickKill} onClickSpawn={onClickSpawn} />)} </div>
    );

    return (
        <div id="App">
            <div className='grid'>
                {gameStateGridBody}
            </div>
            <div className='controller'>
                <AppBar color="primary" sx={{ top: 'auto', bottom: 0 }}>
                    <Toolbar>
                        <IconButton color="inherit">
                            {!isPlaying && <PlayArrow onClick={() => setIsPlaying(true)} />}
                            {isPlaying && <Pause onClick={() => setIsPlaying(false)} />}
                        </IconButton>
                        <IconButton color="inherit">
                            <RestartAlt onClick={() => clearCurrentGameState()} />
                        </IconButton>
                    </Toolbar>
                </AppBar>
                <div>
                    <p> {tick} </p>
                </div>
            </div>
        </div>
    )
}

export default App
