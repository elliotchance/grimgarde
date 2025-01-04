package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

type ItemClass string

const (
	ItemClassHelmet = ItemClass("Helmet")
	ItemClassChest  = ItemClass("Chest")
	ItemClassArms   = ItemClass("Arms")
	ItemClassGloves = ItemClass("Gloves")
	ItemClassBelt   = ItemClass("Belt")
	ItemClassPants  = ItemClass("Pants")
	ItemClassBoots  = ItemClass("Boots")
	ItemClassWeapon = ItemClass("Weapon")
	ItemClassRing   = ItemClass("Ring")
	ItemClassAmulet = ItemClass("Amulet")
)

type AttackRating struct {
	Name             string
	AttacksPerSecond float64
}

var (
	AttackRatingVerySlow = AttackRating{"Very Slow", 0.5}
	AttackRatingSlow     = AttackRating{"Slow", 0.75}
	AttackRatingNormal   = AttackRating{"Normal", 1}
	AttackRatingFast     = AttackRating{"Fast", 1.25}
	AttackRatingVeryFast = AttackRating{"Very Fast", 1.5}
)

type ItemType struct {
	Name         string
	Class        ItemClass
	AttackRating AttackRating // Weapons only
}

var (
	Cap    = ItemType{Name: "Cap", Class: ItemClassHelmet}
	Helmet = ItemType{Name: "Helmet", Class: ItemClassHelmet}

	HideArmor = ItemType{Name: "Hide Armor", Class: ItemClassChest}
	ChainMail = ItemType{Name: "Chain Mail", Class: ItemClassChest}

	LeatherSleeves = ItemType{Name: "Leather Sleeves", Class: ItemClassArms}

	LeatherGloves = ItemType{Name: "Leather Gloves", Class: ItemClassGloves}

	Belt = ItemType{Name: "Belt", Class: ItemClassBelt}

	LeatherPants = ItemType{Name: "Leather Pants", Class: ItemClassPants}

	LeatherBoots = ItemType{Name: "Leather Boots", Class: ItemClassBoots}

	Sword = ItemType{Name: "Sword", Class: ItemClassWeapon, AttackRating: AttackRatingNormal}
	Axe   = ItemType{Name: "Axe", Class: ItemClassWeapon, AttackRating: AttackRatingSlow}
	Mace  = ItemType{Name: "Mace", Class: ItemClassWeapon, AttackRating: AttackRatingSlow}

	Amulet = ItemType{Name: "Amulet", Class: ItemClassAmulet}

	Ring = ItemType{Name: "Ring", Class: ItemClassRing}
)

type Item struct {
	Name         string
	Type         ItemType
	AttackRating AttackRating // Weapons only
	Stats        ItemStats
}

func (item Item) DamagePerHit() int {
	s := item.Stats
	return s.Get(PhysicalDamage) + s.Get(ColdDamage) + s.Get(FireDamage) +
		s.Get(PoisonDamage) + s.Get(LightningDamage)
}

func (item Item) DamagePerSecond() int {
	damage := float64(item.DamagePerHit())
	damage *= item.AttackRating.AttacksPerSecond
	return int(damage)
}

func (item *Item) Lines() []string {
	var lines []string
	if len(item.Name) >= 25 {
		parts := strings.Split(item.Name, " of ")
		lines = append(lines, parts[0], "of "+parts[1])
	} else {
		lines = append(lines, item.Name, "")
	}
	if item.AttackRating.Name != "" {
		lines = append(lines, fmt.Sprintf("%d DPS (%s Attack Rating)", item.DamagePerSecond(), item.AttackRating.Name))
	}
	lines = append(lines, "")
	for _, stat := range item.Stats {
		lines = append(lines, stat.String())
	}
	return lines
}

func (item Item) String() string {
	return strings.Join(item.Lines(), "\n")
}

type ItemStatType struct {
	Name     string
	Min, Max int
}

var (
	// Damage
	PhysicalDamage  = ItemStatType{"Physical Damage", 1, 100}
	ColdDamage      = ItemStatType{"Cold Damage", 1, 100}
	FireDamage      = ItemStatType{"Fire Damage", 1, 100}
	PoisonDamage    = ItemStatType{"Poison Damage", 1, 100}
	LightningDamage = ItemStatType{"Lightning Damage", 1, 100}
	ElementalDamage = ItemStatType{"Elemental Damage", 1, 100}

	// Resistances (absorbs Damage)
	PhysicalResistance  = ItemStatType{"Physical Resistance", 0, 0}
	ColdResistance      = ItemStatType{"Cold Resistance", 0, 0}
	FireResistance      = ItemStatType{"Fire Resistance", 0, 0}
	PoisonResistance    = ItemStatType{"Poison Resistance", 0, 0}
	LightningResistance = ItemStatType{"Lightning Resistance", 0, 0}
	ElementalResistance = ItemStatType{"Elemental Resistance", 1, 100}
	AllResistance       = ItemStatType{"All Resistance", 1, 100}

	// Attack
	AttackSpeed    = ItemStatType{"Attack Speed", 90, 110}
	CriticalChance = ItemStatType{"Critical Chance", 90, 110}
	CriticalDamage = ItemStatType{"Critical Damage", 90, 110}

	// Avoid (opposite of Attack)
	DodgeChance   = ItemStatType{"Dodge Chance", 90, 110}
	ReflectChance = ItemStatType{"Reflect Chance", 90, 110}
	ReflectDamage = ItemStatType{"Reflect Damage", 90, 110}

	// Player Stats
	Life         = ItemStatType{"Life", 0, 0}
	Mana         = ItemStatType{"Mana", 0, 0}
	Strength     = ItemStatType{"Strength", 0, 0}
	Intelligence = ItemStatType{"Intelligence", 0, 0}
	Dexterity    = ItemStatType{"Dexterity", 0, 0}
	Vitality     = ItemStatType{"Vitality", 0, 0}
	AllStats     = ItemStatType{"All Stats", 0, 0}

	// CooldownReduction = ItemStatType{"Cooldown Reduction", 0, 0}

	// life, mana, on kill, per second
	// strength, intelligence, etc
	// provide skills
	// cooldown reduction
	// weapon mastery?
)

var ItemStatTypes = []ItemStatType{
	PhysicalDamage,
	ColdDamage,
	FireDamage,
	PoisonDamage,
	LightningDamage,
	AttackSpeed,
}

type ItemStat struct {
	Value int
	Type  ItemStatType
}

func (s ItemStat) String() string {
	return fmt.Sprintf("+%d %s", s.Value, s.Type.Name)
}

type ItemStats []ItemStat

func (stats ItemStats) Copy() ItemStats {
	return append(ItemStats{}, stats...)
}

func (stats ItemStats) Append(newStats ...ItemStat) ItemStats {
	result := stats.Copy()
	for _, newStat := range newStats {
		if newStat.Value != 0 {
			found := false
			for i, existing := range result {
				if existing.Type == newStat.Type {
					result[i].Value += newStat.Value
					found = true
					break
				}
			}
			if !found {
				result = append(result, newStat)
			}
		}
	}
	return result
}

func (stats ItemStats) Scale(factor float64, statsToScale ...ItemStatType) ItemStats {
	result := stats.Copy()
	for _, statToScale := range statsToScale {
		for i, existing := range result {
			if existing.Type == statToScale {
				result[i].Value = int(float64(result[i].Value) * factor)
				break
			}
		}
	}
	return result
}

func (stats ItemStats) Get(statType ItemStatType) int {
	for _, stat := range stats {
		if stat.Type == statType {
			return stat.Value
		}
	}

	return 0
}

func randomStatType() ItemStatType {
	return ItemStatTypes[rand.Int31n(int32(len(ItemStatTypes)))]
}

func rng(min, max int) int {
	return min + int(rand.Int31n(int32(max-min)))
	// return min + int64(multiplier(rng)*float64(max-min))
}

func multiplier(x float64) float64 {
	// x is 0 to 1, adjust to -2 to 2
	x2 := (x * 4) - 2

	return math.Erfc(-x2) / 2
}

func GenerateRarity() (string, int) {
	values := map[string]int{
		"":          2,
		"Quality":   3,
		"Rare":      4,
		"Legendary": 5,
	}
	for name, stats := range values {
		return name, stats
	}
	return "", 0 // impossible
}

func GeneratePrefix() (string, ItemStat) {
	return pickMap(map[string]ItemStat{
		"Brutal":    {5, PhysicalDamage},
		"Chilling":  {5, ColdDamage},
		"Scorching": {5, FireDamage},
		"Poisoning": {5, PoisonDamage},
		"Sparking":  {5, LightningDamage},
	})
}

func GenerateSuffix() (string, ItemStat) {
	return pickMap(map[string]ItemStat{
		"Violence":    {5, PhysicalDamage},
		"Freezing":    {5, ColdDamage},
		"Burning":     {5, FireDamage},
		"Poison":      {5, PoisonDamage},
		"Electricity": {5, LightningDamage},
	})
}

func GenerateItem(itemType ItemType) *Item {
	rarity, totalStats := GenerateRarity()
	prefix, prefixStat := GeneratePrefix()
	suffix, suffixStat := GenerateSuffix()

	// Don't allow an item to have two of the same stats.
	if prefixStat.Type == suffixStat.Type {
		return GenerateItem(itemType)
	}

	item := &Item{
		Name:         strings.TrimSpace(fmt.Sprintf("%s %s %s of %s", rarity, prefix, itemType.Name, suffix)),
		Type:         itemType,
		AttackRating: itemType.AttackRating,
		Stats:        []ItemStat{prefixStat, suffixStat},
	}

	seen := map[string]struct{}{
		prefixStat.Type.Name: {},
		suffixStat.Type.Name: {},
	}
	for len(item.Stats) < totalStats {
		typ := randomStatType()
		if _, ok := seen[typ.Name]; ok {
			continue
		}
		seen[typ.Name] = struct{}{}
		value := rng(typ.Min, typ.Max)
		item.Stats = append(item.Stats, ItemStat{value, typ})
	}

	return item
}

func pickMap[K comparable, V any](values map[K]V) (K, V) {
	for k, v := range values {
		return k, v
	}

	// impossible
	var emptyK K
	var emptyV V
	return emptyK, emptyV
}

func pickSlice[V any](values []V) V {
	return values[rand.Int()%len(values)]
}

func GenerateHelmet() *Item {
	return GenerateItem(pickSlice([]ItemType{
		Cap,
		Helmet,
	}))
}

func GenerateChest() *Item {
	return GenerateItem(pickSlice([]ItemType{
		HideArmor,
		ChainMail,
	}))
}

func GenerateArms() *Item {
	return GenerateItem(pickSlice([]ItemType{
		LeatherSleeves,
	}))
}

func GenerateGloves() *Item {
	return GenerateItem(pickSlice([]ItemType{
		LeatherGloves,
	}))
}

func GenerateBelt() *Item {
	return GenerateItem(pickSlice([]ItemType{
		Belt,
	}))
}

func GeneratePants() *Item {
	return GenerateItem(pickSlice([]ItemType{
		LeatherPants,
	}))
}

func GenerateBoots() *Item {
	return GenerateItem(pickSlice([]ItemType{
		LeatherBoots,
	}))
}

func GenerateWeapon() *Item {
	return GenerateItem(pickSlice([]ItemType{
		Sword,
		Axe,
		Mace,
	}))
}

func GenerateAmulet() *Item {
	return GenerateItem(pickSlice([]ItemType{
		Amulet,
	}))
}

func GenerateRing() *Item {
	return GenerateItem(pickSlice([]ItemType{
		Ring,
	}))
}

func GenerateAnyItem() *Item {
	return pickSlice([]func() *Item{
		GenerateAmulet,
		GenerateArms,
		GenerateBelt,
		GenerateBoots,
		GenerateChest,
		GenerateGloves,
		GenerateHelmet,
		GeneratePants,
		GenerateRing,
		GenerateWeapon,
	})()
}
