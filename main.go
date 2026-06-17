package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	Version = "0.1.0-nexus"
	Banner  = `
    _   _                     
   | \ | | _____  ___   _ ___ 
   |  \| |/ _ \ \/ / | | / __|
   | |\  |  __/>  <| |_| \__ \
   |_| \_|\___/_/\_\__,_|___/

   Go-Nexus Orchestrator v%s

   Unified starter for mesh + blockchain + AI swarms + prototypes
`
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "start":
		startCmd := flag.NewFlagSet("start", flag.ExitOnError)
		component := startCmd.String("component", "all", "all | mesh | blockchain | ai | prototypes | grok")
		dryRun := startCmd.Bool("dry-run", true, "Show plan and commands without executing (safe default)")
		force := startCmd.Bool("force", false, "Execute even if some checks fail (use with extreme caution)")
		startCmd.Parse(os.Args[2:])
		runStart(*component, *dryRun, *force)

	case "doctor", "check":
		runDoctor()

	case "version", "--version", "-v":
		fmt.Printf("go-nexus %s\n", Version)

	case "help", "--help", "-h":
		printHelp()

	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printHelp()
	}
}

func printHelp() {
	fmt.Printf(Banner, Version)
	fmt.Println(`Usage:
  go-nexus [command] [flags]

Commands:
  start      Start Nexus components (mesh, blockchain, ai, prototypes, grok, all)
  doctor     Run environment health checks and prerequisite report
  version    Show version
  help       Show this help

Flags (for start):
  --component string   Component to start (default "all")
  --dry-run bool       Print commands & rationale without executing (default true)
  --force bool         Force execution despite warnings (default false)

Examples:
  go run main.go doctor
  go run main.go start --component=mesh
  go run main.go start --component=all --dry-run=false --force

Full documentation: see README.md
`)
}

func runDoctor() {
	fmt.Println("\n\ud83e\udda0 Nexus Doctor — Environment Health Check")
	fmt.Println("Checking prerequisites for full Nexus stack (mesh + chain + AI + prototypes)...\n")

	checks := []struct {
		name string
		cmd  string
		args []string
	}{
		{"Go toolchain", "go", []string{"version"}},
		{"Docker", "docker", []string{"--version"}},
		{"Docker Compose", "docker", []string{"compose", "version"}},
		{"Yggdrasil (mesh)", "yggdrasil", []string{"-version"}},
		{"Git", "git", []string{"--version"}},
		{"Rust (Grok Launcher)", "rustc", []string{"--version"}},
		{"Cargo", "cargo", []string{"--version"}},
		{"Python3 (agents/prototypes)", "python3", []string{"--version"}},
	}

	for _, c := range checks {
		if path, err := exec.LookPath(c.cmd); err == nil {
			out, _ := exec.Command(c.cmd, c.args...).CombinedOutput()
			ver := strings.TrimSpace(string(out))
			if ver == "" {
				ver = "(found at " + path + ")"
			}
			fmt.Printf("✅ %-22s %s\n", c.name+":", ver)
		} else {
			fmt.Printf("❌ %-22s NOT FOUND in $PATH\n", c.name+":")
		}
	}

	fmt.Println("\n\ud83d\udcca Recommendations:")
	fmt.Println("  - Install missing components before running full 'start --component=all'.")
	fmt.Println("  - For Yggdrasil: https://yggdrasil-network.github.io/")
	fmt.Println("  - For Grok Launcher (Rust): cargo build in its repo.")
	fmt.Println("  - Docker recommended for containerized agents and services.")
	fmt.Println("\nDoctor complete. Address gaps then re-run.\n")
}

func runStart(component string, dryRun bool, force bool) {
	fmt.Printf("\n\ud83d\ude80 Starting Nexus component: %s  |  Dry-Run: %v  |  Force: %v\n", component, dryRun, force)
	fmt.Println("=" + strings.Repeat("=", 70))
	switch strings.ToLower(component) {
	case "all":
		startMesh(dryRun, force)
		startBlockchain(dryRun, force)
		startAISwarm(dryRun, force)
		startPrototypes(dryRun, force)
		startGrokLauncher(dryRun, force)
	case "mesh":
		startMesh(dryRun, force)
	case "blockchain":
		startBlockchain(dryRun, force)
	case "ai", "swarm":
		startAISwarm(dryRun, force)
	case "prototypes":
		startPrototypes(dryRun, force)
	case "grok", "launcher":
		startGrokLauncher(dryRun, force)
	default:
		fmt.Println("Unknown component. Valid options: all, mesh, blockchain, ai, prototypes, grok")
		return
	}

	fmt.Println("\n\u2705 Startup sequence finished for", component)
	fmt.Println("Monitor logs, verify peer connectivity, and check service status.")
	fmt.Println("Next steps: Run 'go-nexus doctor' periodically and review README.md for scaling & troubleshooting.")
}

func startMesh(dryRun bool, force bool) {
	fmt.Println("\n\ud83d\udce1 [MESH] xMesh / NovaNet / QNET / Yggdrasil / Tenda Nova")
	fmt.Println("   Foundation for decentralized, private, resilient communication.")

	if !checkCmd("yggdrasil") {
		fmt.Println("   \u26a0\ufe0f  Yggdrasil binary not found. Install it first (see doctor output).")
		if !force {
			return
		}
	}

	fmt.Println("   Step 1: Configuration & startup (Yggdrasil)")
	fmt.Println("     → yggdrasil -genconf > ~/.config/yggdrasil.conf   # one-time (review peers)")
	fmt.Println("     → yggdrasil -useconffile ~/.config/yggdrasil.conf")
	fmt.Println("     → # Optional: systemctl --user enable --now yggdrasil")

	fmt.Println("   Nuances: Select diverse bootstrap peers. Tenda Nova — tune channel & TX power (privacy vs range). Docker: use host networking.")
	fmt.Println("   Edge cases: Network partition → multiple bootstraps + auto-reconnect. Privacy: optional Tor/I2P tunneling for sensitive flows.")
	fmt.Println("   Implication: All other layers (chain gossip, AI messages, prototype telemetry) depend on this stable overlay.\n")

	if !dryRun && force {
		fmt.Println("   [LIVE] Attempting status check...")
		if out, err := exec.Command("sh", "-c", "systemctl --user is-active yggdrasil || echo 'inactive/no-systemd-user'").CombinedOutput(); err == nil {
			fmt.Printf("   Status: %s", string(out))
		}
	}
}

func startBlockchain(dryRun bool, force bool) {
	fmt.Println("\n\ud83d\udd17 [BLOCKCHAIN] XCoin / QCoin / QNET runes (Wizard Q)")
	fmt.Println("   Economic layer: incentives, reputation, oracles, governance.")

	fmt.Println("   Step 1: Node sync & startup (placeholder for qcoin-node binary)")
	fmt.Println("     → ./qcoin-node --config ~/.qcoin/config.toml --start   # or future integrated binary")
	fmt.Println("     → Monitor sync progress and peer count")

	fmt.Println("   Nuances & Trade-offs: Token volatility vs utility. Rune system enables on-chain automation. Align rewards with mesh uptime & prototype data quality.")
	fmt.Println("   Edge: Reorgs/finality — design tolerant agent logic. Regulatory (MiCA, US) — leverage Delaware C-Corp structure.")
	fmt.Println("   Implication: Provides economic flywheel for the entire Nexus stack and long-term sustainability.\n")
}

func startAISwarm(dryRun bool, force bool) {
	fmt.Println("\n\ud83e\udd16 [AI SWARM] Self-improving emotional agents + Grok integration")
	fmt.Println("   Intelligence & autonomy layer running over mesh.")

	fmt.Println("   Step 1: Launch agent processes / Grok Launcher bridge")
	fmt.Println("     → # Future: go run agent-swarm/main.go or cargo run (Grok Launcher Rust)")
	fmt.Println("     → # Swarm coordination via Yggdrasil pub/sub channels")

	fmt.Println("   Nuances: Prompt engineering + reputation-weighted consensus to prevent drift/hallucination. Emotional models (Ara) require affective logging + stability guardrails.")
	fmt.Println("   Edge: Agent divergence — human review gates or on-chain proposals for self-modification.")
	fmt.Println("   Implication: Enables autonomous operation, creative generation, and recursive improvement of the whole ecosystem.\n")
}

func startPrototypes(dryRun bool, force bool) {
	fmt.Println("\n\ud83d\udee0️  [PROTOTYPES] Soilnova / Vista Nova / York Autotype / Lumia")
	fmt.Println("   Physical-digital bridge: sensors, actuators, visualization, automation.")

	fmt.Println("   Step 1: Hardware bring-up & data pipelines")
	fmt.Println("     → # Soilnova: I2C sensors on Raspberry Pi / Arduino → signed readings to mesh/oracle")
	fmt.Println("     → # Vista Nova / Lumia: low-power display/lighting scenes controlled via mesh")
	fmt.Println("     → # York Autotype: event-driven automation tied to on-chain or agent triggers")

	fmt.Println("   Nuances: Calibration, secure boot, power management, sensor drift compensation.")
	fmt.Println("   Edge cases: Hardware failure → watchdog + last-known-good state recovery in orchestrator.")
	fmt.Println("   Implication: Grounds abstract stack in real-world data and action — oracles, physical feedback loops.\n")
}

func startGrokLauncher(dryRun bool, force bool) {
	fmt.Println("\n\ud83d\ude80 [GROK LAUNCHER] Rust + egui prototyping environment")
	fmt.Println("   High-performance local UI and inference prototyping tool.")

	fmt.Println("   Step 1: Build & launch (separate repo integration)")
	fmt.Println("     → cd ../grok-launcher && cargo build --release && ./target/release/grok-launcher")
	fmt.Println("     → # Or exec from PATH if installed globally")

	fmt.Println("   Nuances: Use for rapid UI mocks, local model testing, and as visual frontend to Nexus state.")
	fmt.Println("   Integration: Future FFI or IPC bridge from this Go orchestrator for unified control plane.")
	fmt.Println("   Implication: Accelerates development velocity and provides delightful human-in-the-loop interface.\n")
}

func checkCmd(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
