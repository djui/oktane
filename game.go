package main

type Game struct {
	config  Config
	modules []Module
}

type Config struct {
	noColor bool
	noANSI  bool
}

func NewGame(config Config) *Game {
	return &Game{
		config:  config,
		modules: pickModules(),
	}
}

func pickModules() []Module {
	modules := []Module{
		// front
		&KeypadsModule{},
		&TimeModule{},
		&MemoryModule{},
		&MazeModule{},
		&ButtonModule{},
		&WiresModule{},
		// back
		&ComplicatedWiresModule{},
		&WhosOnFirstModule{},
		&SimonSaysModule{},
		&MorseCodeModule{},
		&WireSequencesModule{},
		&PasswordModule{},
		//&CapacitorDischargeModule{},
		//&VentingGasModule{},
		//&KnobModule{},
	}

	modules[0].solve(true)
	modules[2].solve(true)
	modules[4].solve(true)

	return modules
}
