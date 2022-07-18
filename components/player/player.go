package player

type Player struct {
	name string
	hit uint
	miss uint
}

func New(name string) *Player {
	var playerTest = &Player{
		name: name,
		hit:0,
		miss:0,
	}
	return playerTest //pointer to player
}
func (p*Player) GetName() (string){
	return p.name
}
func (p*Player) GetHit() (uint){
	return p.hit
}
func (p*Player) GetMiss() (uint){
	return p.miss
}
func (p*Player) IncrementHit(){
	 p.hit++
}

func (p*Player) IncrementMiss(){
	 p.miss++
}
