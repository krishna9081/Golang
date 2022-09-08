
type Node struct {
    n *TreeNode
    vlevel int
    
}

func verticalOrder(root *TreeNode) [][]int {
    
    //Edge case check
    if root == nil{
       return nil 
    }
    
    //Key = level and Value is lists of node values at that level
    
    m := map[int][]int{}
    
    minvlevel:= 0
    //Intialize 
    rootNode := Node{n:root,vlevel:0,}
    
    //Now we have to start at the root node and traverse down the line and capture level information in the map that we created.
    //We have list of Nodes for BFS and add the root Node
    nodeList := []*Node {&rootNode}
    
    for len(nodeList) > 0 {
        
        // For Iteration take another variable list of nodes
        var nextNodeList []*Node
        //At a specific level iterate over the list of node
        for _,node:=range nodeList{
            currLevel,currNode := node.vlevel,node.n
            m[currLevel]=append(m[currLevel],currNode.Val)
            
            
            if currNode.Left !=nil{
                leftNode := Node{n:currNode.Left,vlevel:currLevel-1,}
                if minvlevel > currLevel-1 {minvlevel=currLevel-1}
                nextNodeList = append(nextNodeList,&leftNode)
            }
            
            if currNode.Right !=nil{
                rightNode := Node{n:currNode.Right,vlevel:currLevel+1,}
                nextNodeList = append(nextNodeList,&rightNode)
            }         
           
        }
         
        nodeList = nextNodeList  
    }
    //Now the map gets populated with levels and Node values
        //Iterate over the map and append the items to the list  
    res := make([][]int, len(m))
    
    for i, v := range m {
        if val;ok:=
    }
    
    return res
       
}



