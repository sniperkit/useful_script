let g:go_highlight_functions = 1
let g:go_highlight_methods = 1
let g:go_highlight_structs = 1
let g:go_highlight_interfaces = 1
let g:go_highlight_operators = 1
let g:go_highlight_build_constraints = 1

let g:spf13_keep_trailing_whitespace = 1
let g:cscope_silent=1

filetype off
set rtp+=~/.vim/bundle/Vundle.vim

let mapleader="."

call vundle#begin()
Bundle 'brookhong/cscope.vim'
Bundle 'Yggdroot/indentLine'
Bundle 'vim-scripts/grep.vim'
Bundle 'vim-scripts/taglist.vim'
Bundle 'scrooloose/nerdtree'
Bundle 'airblade/vim-gitgutter'
"Bundle 'lamproae/tComment.vim'
Bundle 'Blackrush/vim-gocode'
Bundle 'lamproae/vim-go'
Bundle 'tpope/vim-markdown'
"Bundle 'vim-airline/vim-airline'
"Bundle 'vim-airline/vim-airline-themes'
"Bundle 'Valloric/YouCompleteMe'
Bundle 'NLKNguyen/papercolor-theme'
Bundle 'vim-scripts/mru.vim'
Bundle 'vim-scripts/desert256.vim'
Bundle 'vim-syntastic/syntastic'
call vundle#end()
filetype plugin indent on

" s: Find this C symbol
nnoremap  <leader>fs :call CscopeFind('s', expand('<cword>'))<CR>
" g: Find this definition
nnoremap  <leader>fg :call CscopeFind('g', expand('<cword>'))<CR>
" " d: Find functions called by this function
nnoremap  <leader>fd :call CscopeFind('d', expand('<cword>'))<CR>
" " c: Find functions calling this function
nnoremap  <leader>fc :call CscopeFind('c', expand('<cword>'))<CR>
" " t: Find this text string
nnoremap  <leader>ft :call CscopeFind('t', expand('<cword>'))<CR>
" " e: Find this egrep pattern
nnoremap  <leader>fe :call CscopeFind('e', expand('<cword>'))<CR>
" " f: Find this file
nnoremap  <leader>ff :call CscopeFind('f', expand('<cword>'))<CR>
" " i: Find files #including this file
nnoremap  <leader>fi :call CscopeFind('i', expand('<cword>'))<CR

nnoremap <leader>fa :call CscopeFindInteractive(expand('<cword>'))<CR>
nnoremap <leader>l :call ToggleLocationList()<CR>

nnoremap <leader>tt :TlistOpen<CR>
nnoremap <leader>nt :NERDTree<CR>

nmap <C-n> :cnext<CR>
nmap <C-p> :cprev<CR>

nnoremap <leader>rg :Rgrep<CR>
nnoremap <leader>q :q<CR>
nnoremap <leader>w :w<CR>
nnoremap <leader>wq :wq<CR>
nnoremap <leader>g :<C-]><CR>

"set lines=60
set nolist
set noexpandtab
set tabstop=8
set linespace=8
set shiftwidth=4
"set statusline=%F\ %m%h%r%<%=\ [%{&ff},%{$fenc}]\ [%Y]\ [%l,%v]\ [%L]\ [%p%%]
set statusline=%F\ [%{&ff},%{$fenc}]\ [%Y]\ [%l,%v]\ [%L]\ [%p%%]
set ruler
syntax on
set hlsearch
set cindent
set nu
set laststatus=2
set background=dark

let Grep_Default_Filelist='*.[chS]'
let Grep_Default_Options='-i --color=auto'
let Grep_Skip_Dirs='.svn'

let g:indentLine_char='|'
let g:indentLine_color_term=239
"let g:indentLine_color_gui="#A4E57E"
let g:indentLine_color_dark=1
let g:indentLine_enabled=1
let Tlist_Use_Right_Window=1
"let g:airline#extensions#tabline#enabled = 1
"let g:airline#extensions#tabline#left_sep = ' '
let g:airline#extensions#tabline#left_alt_sep = '|'
let MRU_Max_Menu_Entries = 20

"colorscheme koehler
"colorscheme evening
"colorscheme morning
"colorscheme pablo
"colorscheme ron
"colorscheme shine
"colorscheme torte
"colorscheme zeliner
"colorscheme hybrid
"colorscheme jellybeans
"colorscheme PaperColor
"colorscheme desert
colorscheme desert256

"autocmd vimenter * NERDTree
"autocmd vimenter *  TlistOpen
autocmd StdinReadPre * let s:std_in=1
autocmd VimEnter * if argc() == 0 && !exists("s:std_in") | NERDTree | endif
autocmd StdinReadPre * let s:std_in=1
autocmd VimEnter * if argc() == 1 && isdirectory(argv()[0]) && !exists("s:std_in") | exe 'NERDTree' argv()[0] | wincmd p | ene | endif
autocmd bufenter * if (winnr("$") == 1 && exists("b:NERDTree") && b:NERDTree.isTabTree()) | q | endif
map <C-q> :NERDTreeToggle<CR>

set updatetime=250

ab epd extern void printd(const char * fmt, , .);
ab pd printd("[%s][%d]-------------------\n", __func__, __LINE__);
ab pk printk(KERN_EMERG "[%s][%d]-------------------\n", __func__, __LINE__);
ab .. , 
